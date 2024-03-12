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


type xxxRequest struct {

}

type xxxResponse struct {

}



// 获取登录页面
func (h *UserHandler) GetLogin(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}	


// 用户登录（通过账号，用户名和密码）
func (h *UserHandler) DoLoginByAccount(c *gin.Context) {

}


// 用户登录（通过发送邮件验证码）
func (h *UserHandler) DoLoginByEmailOTP(c *gin.Context) {

}

// 用户登录（通过第三方授权）
func (h *UserHandler) DoLoginByGithub(c *gin.Context) {

}

