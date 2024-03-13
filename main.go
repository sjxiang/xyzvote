package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"xyzvote/api"
	"xyzvote/consts"
	"xyzvote/db"
)


type Server struct {
	
}


func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	database, err := gorm.Open(mysql.Open(consts.MySQLDefaultDSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	
	var (
		captchaDigitDriver = base64Captcha.DriverDigit{  // 数字
			Height:   consts.CaptchaHeight,  
			Width:    consts.CaptchaWidth, 
			Length:   consts.CaptchaLength,   
			MaxSkew:  consts.CaptchaMaxSkew,   
			DotCount: consts.CaptchaDotCount,   
		}
		captchaStore       = base64Captcha.DefaultMemStore  // 相关数据（e.g. 验证码）存储在内存中
		captcha            = base64Captcha.NewCaptcha(&captchaDigitDriver, captchaStore)
		
		userStore          = db.NewMySQLUserStore(database)
		// voteStore          = db.NewMySQLVoteStore(database)
		
		userHandler        = api.NewUserHandler(userStore, captcha)
		// voteHandler        = api.NewVoteHandler(voteStore, userStore)
		
		app                = gin.Default()
		apiv1              = app.Group("/api/v1")
	)

	app.LoadHTMLGlob("api/views/*")

	apiv1.GET("/otp/gen", userHandler.GenerateImageOTP)
	apiv1.POST("/otp/verify", userHandler.VerifyImageOTP)
	apiv1.POST("/user/register", userHandler.Register)	
	apiv1.POST("/user/login", userHandler.LoginByAccount)

	apiv1.Use(userHandler.Check())

	apiv1.GET("/user/me", userHandler.Me)
	apiv1.GET("/user/admin", userHandler.Admin)




	app.Run(":8080")
}




// func runGinServer(config util.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create server")
// 	}

// 	err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot start server")
// 	}
// }

// dependency injection，依赖注入
