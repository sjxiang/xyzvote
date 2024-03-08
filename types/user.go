package types

import (
	"fmt"
	"time"
)


const (
	bcryptCost      = 12
	minUserNameLen  = 2
	minPasswordLen  = 7
)


type LoginUserParams struct {
	Username      string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

func (params LoginUserParams) Validate() map[string]string {
	errors := map[string]string{}
	
	if len(params.Username) < minUserNameLen {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUserNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}

	return errors
}


type User struct {
    Id                 int64     `json:"id,omitempty" gorm:"column:id;primary_key"`
    UserId             int64     `json:"user_id"      gorm:"column:user_id"`
    Username           string    `json:"username"     gorm:"column:username"`
    EncryptedPassword  string    `json:"-"            gorm:"column:password"`
    Email              string    `json:"email"        gorm:"column:email"`
    Gender             int8      `json:"gender"       gorm:"column:gender"`
    CreatedAt          time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt          time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

