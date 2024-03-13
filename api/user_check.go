package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 校验用户身份

// cookie 校验
func (h *UserHandler) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := c.Cookie("credentials")
		if err != nil || creds == "" {
			c.Redirect(http.StatusPermanentRedirect, "/api/v1/user/login")
		}

		c.Set("creds", creds)
		c.Next()
	}
}


// jwt 校验
func (h *UserHandler) AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
