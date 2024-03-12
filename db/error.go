package db

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrRecordNoFound      = errors.New("db: no matching record found")
	ErrInvalidCredentials = errors.New("db: invalid credentials")
	ErrDuplicateEmail     = errors.New("db: duplicate email")
	ErrDuplicateUsername  = errors.New("db: duplicate username")
	ErrDuplicateUserId    = errors.New("db: duplicate user_id")
)


// 抽离出 err number
func ErrorNumber(err error) int {
	var mySQLErr *mysql.MySQLError

	if errors.As(err, &mySQLErr) {
		return int(mySQLErr.Number)
	}
	return 0
}
