package api

import (
	"context"
	"errors"
	"net/http"
	"xyzvote/db"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// 问卷调查首页，榜单
func (h *VoteHandler) RankList(c *gin.Context) {
	forms, err := h.voteStore.ListForms(context.TODO())
	if err != nil {
		if errors.Is(err, db.ErrRecordNoFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "表单为空",
			})
		}

		log.Error().Err(err).Msg("db cannot list forms")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"form": forms,
	})
}


