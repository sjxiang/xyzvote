package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"xyzvote/db"
	"xyzvote/types"
)



type UserHandler struct {
	// db
	userStore db.UserStore
}


func NewUserHandler(store db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: store,
	}
}

func (h *UserHandler) GetLogin(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}	


func (h *UserHandler) DoLogin(c *gin.Context)  {
	var params types.LoginUserParams
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}
	
	result, err := h.userStore.GetUserByUsername(context.Background(), params.Username)
	if err != nil {
		// 检查 ErrRecordNotFound 错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "账号密码错误",
			})
			return
		}

		// 数据库查询失败
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "查询失败",
		})
		return
	}

	if result.EncryptedPassword != params.Password {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "账号密码错误",
		})
		return
	}

	oneHour := int(time.Second * 60 * 60)
	c.SetCookie("credentials", result.Username, oneHour, "/", "", true, false)

	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}	


func (h *UserHandler) Admin(c *gin.Context)  {
	result, err := h.userStore.GetUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusBadGateway,  gin.H{
			"msg": "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *UserHandler) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := c.Cookie("credentials")
		if err != nil || creds == "" {
			c.Redirect(http.StatusPermanentRedirect, "/api/v1/login")
		}

		c.Set("creds", creds)
		c.Next()
	}
}



func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("credentials", "", 3600, "/", "", true, false)
	c.Redirect(http.StatusPermanentRedirect, "/api/v1/login")
}