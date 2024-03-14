package api

import (
	"context"
	"errors"
	"net/http"
	"xyzvote/db"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type GetVoteRequest struct {
	FormID int64 `form:"form_id" url:"form_id" binding:"required,min=1"`
}

// 查询投票详情
func (h *VoteHandler) GetVoteInfo(c *gin.Context) {
	var req GetVoteRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	form, err := h.voteStore.GetForm(context.TODO(), req.FormID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNoFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "无效的表单编号",
			})
			return
		}

		log.Error().Err(err).Msg("db cannot get form")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	options, err := h.voteStore.GetOptionByFormId(context.TODO(), req.FormID)
	log.Error().Err(err).Msg("what happened")
	
	if err != nil && !errors.Is(err, db.ErrRecordNoFound){
		log.Error().Err(err).Msg("db cannot get option by form_id")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	
	c.HTML(http.StatusOK, "vote.html", gin.H{
		"form":    form,
		"options": options,
	})

}







	/*

	GET /api/v1/form?id=1

	换种写法，参考

	var (
		idStr = c.Query("id")
		id  int64
		err error
	)
	id, err = strconv.ParseInt(idStr, 10, 64)
	
	
	换种写法，参考
	GET /api/v1/form/:id

	var id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid vote ID",
		})
		return
	}

	*/

	