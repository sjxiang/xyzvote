package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"xyzvote/db"
)


type meResponse struct {
	UserId             string    `json:"user_id"      gorm:"column:user_id"`
    Username           string    `json:"username"     gorm:"column:username"`
    Email              string    `json:"email"        gorm:"column:email"`
    Gender             int8      `json:"gender"       gorm:"column:gender"`
}

// 个人资料
func (h *UserHandler) Me(c *gin.Context) {
	// c 上下文
	username := c.MustGet("creds").(string)

	user, err := h.userStore.GetUser(context.TODO(), username)
	if err != nil {
		if errors.Is(err, db.ErrInvalidCredentials) {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "账号密码错误",
			})
			return
		}

		log.Error().Err(err).Msg("db cannot get user")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	var rsp meResponse
	
	rsp.UserId    = user.UserId
	rsp.Username  = user.Username
	rsp.Email     = user.Email
	rsp.Gender    = user.Gender

	c.JSON(http.StatusOK, gin.H{
		"data": rsp,
	})
}