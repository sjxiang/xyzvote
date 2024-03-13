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
	Username  string `json:"username" form:"username"`
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
    UserId             string    `json:"user_id"      gorm:"column:user_id"`
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


type RegisterUserParams struct {
	Username         string    `json:"username" binding:"required,min=8"`         // 用户名不为空，最少 8 位
    Password         string    `json:"password" binding:"required,gte=8,lte=16"`  // 密码不为空，8 ~ 16 位，不能为纯数字
	ConfirmPassword  string    `json:"confirm_password" binding:"required,gte=8,lte=16"`
	Email            string    `json:"email" binding:"required,email"`
}

type VerifyOTPParams struct {
	CaptchaId string `json:"captcha_id"`
	Data      string `json:"data"`
} 