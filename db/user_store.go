package db

import (
	"context"

	"gorm.io/gorm"

	"xyzvote/types"
)

type Map map[string]any

type UserStore interface {
	GetUserByUsername(ctx context.Context, username string) (*types.User, error)
	// GetUserByEmail(context.Context, string) (*types.User, error)
	// GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) error
	// DeleteUser(context.Context, string) error
	// UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type MySQLUserStore struct {
	storage *gorm.DB
}

func NewMySQLUserStore(storage *gorm.DB) *MySQLUserStore {
	return &MySQLUserStore{
		storage: storage,
	}
}

func (s *MySQLUserStore) GetUserByUsername(ctx context.Context, username string) (*types.User, error) {
	var (
		item types.User
		err  error
	)
	/*
		First v.s. Find

		First
			SELECT * FROM `user` ORDER BY `id` LIMIT 1;
		
			限定词较多
				
		Find
			SELECT * FROM `user` 
			
			全表扫描

	*/

	// record not found
	// SELECT * FROM `user` WHERE username = 'szf199706' ORDER BY `user`.`id` LIMIT 1
	err = s.storage.Table("user").Where("username = ?", username).First(&item).Error  
	return &item, err
}

func (s *MySQLUserStore) GetUsers(context.Context) ([]*types.User, error) {
	var (
		items []*types.User
		err   error 
	) 
	err = s.storage.Table("user").Find(&items).Error
	return items, err
}


func (s *MySQLUserStore) InsertUser(ctx context.Context, user *types.User) error {
	return s.storage.Table("user").Create(&user).Error
}