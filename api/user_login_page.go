package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) LoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}
