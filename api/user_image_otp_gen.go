package api

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// 请求图片 otp
func (h *UserHandler) GenerateImageOTP(c *gin.Context) {
	var (
		id     string
		b64s   string
		answer string
		err    error
	)

	id, b64s, answer, err = h.captcha.Generate()
	if err != nil {
		log.Error().Err(err).Msg("captcha cannot generate otp")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	log.Info().Str("otp", answer).Msg("captcha generate otp")
	
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"captcha_id": id, 
			"data":       b64s,
		},
	})
}

