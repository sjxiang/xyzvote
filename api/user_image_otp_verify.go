package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type verifyImageOTPRequest struct {
	CaptchaId string `json:"captcha_id"`
	Data      string `json:"data"`
} 

// 验证图片 otp
func (h *UserHandler) VerifyImageOTP(c *gin.Context) {
	var req verifyImageOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	// 如果设置 clear 为 true，验证无论成功或失败，otp 都作废
	if ok := h.captcha.Verify(req.CaptchaId, req.Data, true); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "验证成功",
	})
}