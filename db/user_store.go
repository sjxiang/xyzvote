package db

import (
	"gorm.io/gorm"
)


type MySQLUserStore struct {
	database *gorm.DB
}

func NewMySQLUserStore(database *gorm.DB) *MySQLUserStore {
	return &MySQLUserStore{
		database: database,
	}
}
