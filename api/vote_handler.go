package api

import (
	"context"
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"xyzvote/db"
	"xyzvote/types"
)



type VoteHandler struct {
	// db
	voteStore db.VoteStore
	userStore db.UserStore
}


func NewVoteHandler(voteStore db.VoteStore, userStore db.UserStore) *VoteHandler {
	return &VoteHandler{
		voteStore: voteStore,
		userStore: userStore,
	}
}

func (h *VoteHandler) Index(c *gin.Context)  {
	result, err := h.voteStore.GetVotes(context.Background())
	if err != nil {
		c.JSON(http.StatusBadGateway,  gin.H{
			"msg": "db 查询失败",
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusBadGateway,  gin.H{
			"msg": "查询为空",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"vote": result,
	})
}

func (h *VoteHandler) DoVote(c *gin.Context)  {

	var params types.VoteParams 
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	// c 上下文
	username := c.MustGet("creds").(string)

	user, err := h.userStore.GetUser(context.Background(), username)
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
			"msg": "db 查询失败",
		})
		return
	}

	/*
		防止刷票
	 */
	result, err := h.voteStore.GetUserVoteRecord(context.Background(), user.UserId, params.VoteId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 操作失败",
		})
		return
	}
	if len(result) >= 1 {
		c.JSON(http.StatusConflict, gin.H{
			"msg": "您已经投过票",
		})
		return
	}

	/*
		由于 vote_opt 是对已有选项计数， +1
		vote_opt_user 是新增用户投票记录，可单选或者多选
		故，无冲突
	 */
	err = h.voteStore.InsertUserVoteRecord(context.Background(), user.UserId, params.VoteId, params.VoteOptions)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 操作失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "投票完成",
	})
}





func (h *VoteHandler) GetVote(c *gin.Context)  { 
	/*

	GET /api/v1/vote?id=1

	换种写法，参考

	var (
		idStr = c.Query("id")
		id  int64
		err error
	)
	id, err = strconv.ParseInt(idStr, 10, 64)
	
	
	换种写法，参考
	GET /api/v1/vote/:id

	var id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid vote ID",
		})
		return
	}

	*/
	var params types.GetVoteParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	result1, err := h.voteStore.GetVoteByID(context.Background(), params.ID)
	if err != nil {
		// 检查 ErrRecordNotFound 错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "无相关投票活动",
			})
			return
		}

		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 查询失败",
		})
		return
	}

	result2, err := h.voteStore.GetOptionsByVoteID(context.Background(), params.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 查询失败",
		})
		return
	}
	if len(result2) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "无相关选项",
		})
		return
	}

	c.HTML(http.StatusOK, "vote.html", gin.H{
		"vote":    result1,
		"options": result2,
	})
}
