package db

import (
	"context"

	"gorm.io/gorm"

	"xyzvote/types"
)

type Map map[string]any

type UserStore interface {
	GetUser(ctx context.Context, username string) (*User, error) 
	
	// GetUserByEmail(context.Context, string) (*types.User, error)
	// GetUserByID(context.Context, string) (*types.User, error)
	
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]*User, error)
	// DeleteUser(context.Context, string) error
	// UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type MySQLUserStore struct {
	database *gorm.DB
}

func NewMySQLUserStore(database *gorm.DB) *MySQLUserStore {
	return &MySQLUserStore{
		database: database,
	}
}


// 创建1条
func (s *MySQLUserStore) InsertUser(ctx context.Context, item *types.User) error {
	return s.database.Table("user").Create(&item).Error
}

// 查询1条
func (s *MySQLUserStore) GetUserByName(ctx context.Context, username string) (*types.User, error) {
	var (
		item types.User
		err  error
	)
	/*
		
	First
		SELECT * FROM `user` ORDER BY `id` LIMIT 1;
	
		获取第一条记录（主键升序），查询不到数据则返回 ErrRecordNotFound
		
		特征：限定词较多

	*/

	err = s.database.Table("user").WithContext(ctx).Where("username = ?", username).First(&item).Error  
	return &item, err
}

func (s *MySQLUserStore) GetUserByNameFixBug(ctx context.Context, username string) ([]*types.User, error) {
	var (
		item []*types.User
		err  error
	)
	/*
	
	Find
		SELECT * FROM `user` 

	*/

	err = s.database.Table("user").Debug().WithContext(ctx).Where("username = ?", username).Find(&item).Error  
	return item, err
}

// 查询多条数据
func (s *MySQLUserStore) MGetUsers(context.Context) ([]*types.User, error) {
	var (
		err   error	
		items []*types.User  // 风格统一
	)

	err = s.database.Debug().Table("user").Find(&items).Error
	return items, err
}


// 分页
// limit
func (s *MySQLUserStore) MGetUsersX(context.Context) ([]types.User, error) {
	items := make([]types.User, 0)

	if err := s.database.Debug().Table("user").Find(&items).Error; err != nil {
		// db query fail
		return nil, err
	}

	return items, nil
}

/*	
	var items []*types.User  // 申明变量，但未初始化；即未分配内存，nil

	items = make([]*types.User, 0)  // 分配了内存
	
	*/ 

/*
	gorm First 和 Find，都要传指针，即变量内存地址；
	
	都对 `*`、`**` 做了适配，不管三七二十七，先 & 再说

	gorm Find 对 nil 做了适配，根据变量内存地址 reflect 出原型，再分配一个空切片替代

	*/

