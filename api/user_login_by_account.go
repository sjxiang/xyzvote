package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"xyzvote/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)


const (
	minUserNameLen  = 2
	minPasswordLen  = 8
)


type loginByAccountRequest struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

func (req loginByAccountRequest) Validate() map[string]string {
	errors := map[string]string{}
	
	if len(req.Username) < minUserNameLen {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUserNameLen)
	}
	if len(req.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}

	return errors
}


// 用户登录（通过账号，用户名和密码）
func (h *UserHandler) LoginByAccount(c *gin.Context) {
	var req loginByAccountRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}
	
	user, err := h.userStore.GetUser(context.TODO(), req.Username)
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

	match, err := PasswordMatches(user.EncryptedPassword, req.Password)
	if err != nil {
		log.Error().Err(err).Msg("bcrypt cannot compare password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号密码错误",
		})
		return		
	}
	
	oneHour := int(time.Second * 60 * 60)
	c.SetCookie("credentials", req.Username, oneHour, "/", "", true, false)  

	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

func PasswordMatches(hashedPassword, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

