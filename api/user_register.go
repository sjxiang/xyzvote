package api

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"time"
	"xyzvote/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/rs/zerolog/log"
)



type RegisterUserRequest struct {
	Username         string    `json:"username" binding:"required,min=8"`         // 用户名不为空，最少 8 位
    Password         string    `json:"password" binding:"required,gte=8,lte=16"`  // 密码不为空，8 ~ 16 位，不能为纯数字
	ConfirmPassword  string    `json:"confirm_password" binding:"required,gte=8,lte=16"`
	Email            string    `json:"email" binding:"required,email"`
}

type RegisterUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func newRegisterUserResponse(item *db.User) RegisterUserResponse {
	return RegisterUserResponse{
		Username: item.Username,
		Email:    item.Email,
	}
}


func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}
	
	// 密码一致
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "两次密码不同",
		})
		return
	}

	// 密码不能是纯数字
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码不能为纯数字",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("bcrypt cannot generate hash")
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	arg := db.CreateUserParams{
		UserId:         uuid.New().String(),  // 默认，v4 版本
		Username:       req.Username,
		HashedPassword: string(hashedPassword),
		Email:          req.Email ,
		CreatedAt:      time.Now(),
	}

	user, err := h.userStore.CreateUser(context.TODO(), arg)
	if err != nil {
		if errors.Is(err, db.ErrDuplicateUserId) {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "uuid 重复",  
			})
			return
		}
		if errors.Is(err, db.ErrDuplicateUsername) {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "该用户名已经注册",  
			})
			return
		}
		if errors.Is(err, db.ErrDuplicateEmail) {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "该邮箱已经注册",  
			})
			return
		}

		log.Error().Err(err).Msg("db cannot create user")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	
	rsp := newRegisterUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"msg": "用户注册成功",
		"data": rsp,
	})
}
