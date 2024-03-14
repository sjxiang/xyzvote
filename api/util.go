package api

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}


func GetUUID() string {
	id := uuid.New() // 默认 V4 版本，字符串，不适合做主键
	return id.String()
}


func GetUid() int64 {
	node, _ := snowflake.NewNode(1)  // 雪花算法
	return node.Generate().Int64()
}