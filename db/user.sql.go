package db

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"xyzvote/consts"
)

/*

	db.First(&item)

		1. raw-sql
		SELECT * FROM `user` ORDER BY `id` LIMIT 1;

		获取第一条记录（主键升序），查询不到数据则返回 ErrRecordNotFound

		特征：限定词较多


	db.Find(&items)
		1. raw-sql
		SELECT * FROM `user`



*/

func (s *MySQLUserStore) GetUser(ctx context.Context, username string) (*User, error) {
	
	var item User
	
	if err := s.database.Debug().WithContext(ctx).
		Table(consts.UserTableName).Where("name = ?", username).First(&item).Error; 
		err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}	
	}

	return &item, nil
}


const getUser = `
select * from user where username = ? limit 1
`

// raw sql 优化 
func (s *MySQLUserStore) XGetUser(ctx context.Context, username string) (*User, error) {
	var item User
	
	if err := s.database.Raw(getUser, username).Scan(&item).Error; err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, ErrInvalidCredentials
	}

	return &item, nil
}


type CreateUserParams struct {
	UserId         string    `json:"user_id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at`
}

func (s *MySQLUserStore) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	
	var item User

	item.UserId            = arg.UserId
	item.Username          = arg.Username
	item.EncryptedPassword = arg.HashedPassword
	item.Email             = arg.Email
	item.CreatedAt         = arg.CreatedAt

	if err := s.database.Debug().Table(consts.UserTableName).Create(&item).Error; err != nil {

		var mySQLError *mysql.MySQLError
		
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 {
				switch {
				case strings.Contains(mySQLError.Message, "user.idx_user_id"):
					return nil, ErrDuplicateUserId	
				case strings.Contains(mySQLError.Message, "user.idx_username"):
					return nil, ErrDuplicateUsername
				case strings.Contains(mySQLError.Message, "user.idx_email"):
					return nil, ErrDuplicateEmail
				}
			}
		}
					
		return nil, err
	}

	return &item, nil 
}

func (s *MySQLUserStore) XCreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	panic("implement me!")
}

func (s *MySQLUserStore) CreateUserTx(ctx context.Context, arg CreateUserParams) (User, error) {
	panic("implement me!") 
}



type ListUserParams struct {
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}
	
func (s *MySQLUserStore) ListUser(ctx context.Context, arg ListUserParams) ([]*User, error) {

	items := make([]*User, 0)
	if err := s.database.Debug().Table(consts.UserTableName).Limit(arg.Limit).Offset(arg.Offset).Find(&items).Error; err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrRecordNoFound
	}

	return items, nil 
}


// var (
// 	err   error	
// 	items []*types.User  // 风格统一
// )

// err = s.database.Debug().Table("user").Find(&items).Error
// return items, err

// 游标