package api

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"

	"xyzvote/db"
)


type (
	
	UserHandler struct {
		// db
		userStore db.UserStore
		// cache
		cache     *redis.Client
		// util
		captcha   *base64Captcha.Captcha
	}

	VoteHandler struct {
		// db
		voteStore db.VoteStore
		userStore db.UserStore
	}
	
)

func NewUserHandler(userStore db.UserStore, cache *redis.Client, captcha *base64Captcha.Captcha) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		cache:     cache,
		captcha:   captcha,
	}
}


func NewVoteHandler(voteStore db.VoteStore, userStore db.UserStore) *VoteHandler {
	return &VoteHandler{
		voteStore: voteStore,
		userStore: userStore,
	}
}

