package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mojocn/base64Captcha"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"xyzvote/api"
	"xyzvote/db"
)


type Server struct {
	
}


func main() {

	var (
		dsn           = Env("MysqlDataSourceName", "root:my-secret-pw@tcp(127.0.0.1:3306)/xyz_vote?parseTime=true")
		height, _     = strconv.ParseInt(Env("CaptchaHeight", "80"), 10, 64)
		width, _      = strconv.ParseInt(Env("CaptchaWidth", "240"), 10, 64)
		length, _     = strconv.ParseInt(Env("CaptchaLength", "6"), 10, 64)
		maxSkew, _    = strconv.ParseFloat(Env("CaptchaMaxSkew", "0.7"), 64)
		dotCount, _   = strconv.ParseInt(Env("CaptchaDotCount", "80"), 10, 64)
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	var (
		captchaDigitDriver = base64Captcha.DriverDigit{
			Height:   int(height),  
			Width:    int(width), 
			Length:   int(length),   
			MaxSkew:  maxSkew,   
			DotCount: int(dotCount),   
		}
		captchaStore       = base64Captcha.DefaultMemStore  // 相关数据（e.g. 验证码）存储在内存中
		captcha            = base64Captcha.NewCaptcha(&captchaDigitDriver, captchaStore)
		
		userStore          = db.NewMySQLUserStore(database)
		voteStore          = db.NewMySQLVoteStore(database)
		
		userHandler        = api.NewUserHandler(userStore, captcha)
		voteHandler        = api.NewVoteHandler(voteStore, userStore)
		
		app                = gin.Default()
		apiv1              = app.Group("/api/v1")
	)

	app.LoadHTMLGlob("api/views/*")

	apiv1.GET("/otp/gen", userHandler.GenerateOTP)
	apiv1.POST("/otp/verify", userHandler.VerifyOTP)
	apiv1.GET("/login", userHandler.GetLogin)      // 登录页面
	apiv1.POST("/login", userHandler.DoLogin)      // 登录
	apiv1.POST("/register", userHandler.Register)  // 注册

	apiv1.Use(userHandler.Check())

	apiv1.GET("/index", voteHandler.Index)
	apiv1.GET("/vote", voteHandler.GetVote)
	apiv1.POST("/vote", voteHandler.DoVote)

	apiv1.POST("/vote/add", voteHandler.AddVote)
	apiv1.PUT("/vote/update", voteHandler.EditVote)
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
		log.Println("使用 "+ key + " 缺省值")
		return fallbackValue
	}
	return s
}

// dependency injection，依赖注入
