package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"xyzvote/api"
	"xyzvote/db"
)


func main() {

	dsn := os.Getenv("MySQL_Data_Source_Name")
	storage, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var (
		userStore   = db.NewMySQLUserStore(storage)
		voteStore   = db.NewMySQLVoteStore(storage)
		
		userHandler = api.NewUserHandler(userStore)
		voteHandler = api.NewVoteHandler(voteStore, userStore)
		
		app         = gin.Default()
		apiv1       = app.Group("/api/v1")
	)

	app.LoadHTMLGlob("api/views/*")

	apiv1.GET("/login", userHandler.GetLogin)
	apiv1.POST("/login", userHandler.DoLogin)

	apiv1.Use(userHandler.Check())

	apiv1.GET("/index", voteHandler.Index)
	apiv1.GET("/vote", voteHandler.GetVote)
	apiv1.POST("/vote", voteHandler.DoVote)

	apiv1.POST("/vote/add", voteHandler.AddVote)
	apiv1.PUT("/vote/update", voteHandler.UpdateVote)
	apiv1.DELETE("/vote/del", voteHandler.DelVote)
	
	apiv1.GET("/admin", userHandler.Admin)

	app.Run(":8080")

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func Env(key, fallbackValue string) string {
	s, ok := os.LookupEnv(key)
	if !ok { 
		return fallbackValue
	}
	return s
}

// dependency injection，依赖注入