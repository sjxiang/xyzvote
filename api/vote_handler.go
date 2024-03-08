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
	store db.Store
}


func NewVoteHandler(store db.Store) *VoteHandler {
	return &VoteHandler{
		store: store,
	}
}

func (h *VoteHandler) Index(c *gin.Context)  {
	result, err := h.store.Vote.GetVotes(context.Background())
	if err != nil {
		c.JSON(http.StatusBadGateway,  gin.H{
			"msg": "查询失败",
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

func (h *VoteHandler) GetVoteInfo(c *gin.Context)  { 
	/*

	GET /api/v1/vote?id=1

	换种写法，参考

	var (
		idStr = c.Query("id")
		id  int64
		err error
	)
	id, err = strconv.ParseInt(idStr, 10, 64)
	
	*/
	
	var params types.GetVoteParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	result1, err := h.store.Vote.GetVoteByID(context.Background(), params.ID)
	if err != nil {
		// 检查 ErrRecordNotFound 错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "无相关投票活动",
			})
			return
		}

		// 数据库查询失败
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "查询失败",
		})
		return
	}

	result2, err := h.store.Vote.GetOptionsByID(context.Background(), params.ID)
	if err != nil {
		// 数据库查询失败
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "查询失败",
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


func (h *VoteHandler) DoVote(c *gin.Context)  {
	username := c.MustGet("creds").(string)

	var params types.VoteParams 
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	result, err := h.store.User.GetUserByUsername(context.Background(), username)
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
			"msg": "查询失败",
		})
		return
	}
	err = h.store.Vote.InsertVoteOptions(context.Background(), result.UserId, params.VoteId, params.VoteOptions)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "投票失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "投票完成",
	})
}