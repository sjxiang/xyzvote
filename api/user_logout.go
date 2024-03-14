package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("credentials", "", 3600, "/", "", true, false)
	c.Redirect(http.StatusPermanentRedirect, "/api/v1/user/login")
}