package api

import (
	"context"
	"net/http"
	"errors"
	
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"xyzvote/db"
)

// 后台
type listUserRequest struct {
	PageID   int `json:"page_id" binding:"required,min=1"`
	PageSize int `json:"page_size" binding:"required,min=5,max=10"`
}


type userResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   int8   `json:"gender"`
}


func convertUser(item *db.User) userResponse {
	return userResponse{
		Username: item.Username,
		Email:    item.Email,
		Gender:   item.Gender,
	}
}

func newListUserResponse(items []*db.User) []userResponse {
	
	users := make([]userResponse, 0)
	
	for _, item := range items {
		user := convertUser(item)
		users = append(users, user)
	} 

	return users
}


func (h *UserHandler) Admin(c *gin.Context) {
	var req listUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	arg := db.ListUserParams{
		Limit:   req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := h.userStore.ListUser(context.TODO(), arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNoFound) {
			c.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		log.Error().Err(err).Msg("db cannot list user")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	rsp := newListUserResponse(users)
	c.JSON(http.StatusOK, rsp)
} 
