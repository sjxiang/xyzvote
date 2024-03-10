package api

import (
	// "context"
	// "errors"
	// "log"
	"net/http"
	// "regexp"
	// "strings"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/go-sql-driver/mysql"
	// "github.com/google/uuid"
	// "github.com/mojocn/base64Captcha"
	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"

	// "xyzvote/db"
	// "xyzvote/types"
)


// 登陆页面
func (h *UserHandler) GetLogin(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}	


// 用户登录
