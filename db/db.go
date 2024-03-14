package db

import "gorm.io/gorm"

type (
	
	MySQLVoteStore struct {
		database *gorm.DB
	}
	
	MySQLUserStore struct {
		database *gorm.DB
	}

)

func NewMySQLVoteStore(database *gorm.DB) *MySQLVoteStore {
	return &MySQLVoteStore{
		database: database,
	}
}

func NewMySQLUserStore(database *gorm.DB) *MySQLUserStore {
	return &MySQLUserStore{
		database: database,
	}
}
