package api

import (
	"github.com/mojocn/base64Captcha"
	
	"xyzvote/db"
)


type UserHandler struct {
	// db
	userStore db.UserStore
	// util
	captcha   *base64Captcha.Captcha
}


func NewUserHandler(userStore db.UserStore, captcha *base64Captcha.Captcha) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		captcha:   captcha,
	}
}