package db

import (
	// "context"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func prepareUser(t *testing.T) UserStore {
	var (
		dsn = "root:my-secret-pw@tcp(127.0.0.1:3306)/xyz_vote?parseTime=true"
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		
	})
	if err != nil {
		t.Fatal(err)
	}

	return NewMySQLUserStore(database)
}

func TestXxx(t *testing.T) {
	// userStore := prepareUser(t)
	
	// t.Logf("%+v\n", items)
}