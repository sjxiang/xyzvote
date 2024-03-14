package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"xyzvote/api"
	"xyzvote/consts"
	"xyzvote/db"
)


func main() {
	srv, err := NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = srv.Start(consts.HTTPServerAddres)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}


type Server struct {
	userHandler *api.UserHandler
	voteHandler *api.VoteHandler
	router      *gin.Engine
}

func NewServer() (*Server, error) {
	
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	database, err := gorm.Open(mysql.Open(consts.MySQLDefaultDSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     consts.RedisServerAddr,
		Password: consts.RedisPassword,
		DB:       consts.RedisDatabaseNum,
	})

	_, err = cache.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to cache")
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
		voteStore          = db.NewMySQLVoteStore(database)
		
		userHandler        = api.NewUserHandler(userStore, cache, captcha)
		voteHandler        = api.NewVoteHandler(voteStore, userStore)
	)

	srv := &Server{
		userHandler: userHandler,
		voteHandler: voteHandler,
	}

	srv.setupRouter()
	return srv, nil 
}


func (server *Server) setupRouter() {
	router := gin.Default()
	
	router.LoadHTMLGlob("api/views/*")

	apiv1 := router.Group("/api/v1")

	apiv1.GET("/otp/gen", server.userHandler.GenerateImageOTP)
	apiv1.POST("/otp/verify", server.userHandler.VerifyImageOTP)
	apiv1.POST("/user/register", server.userHandler.Register)	
	apiv1.GET("/user/login", server.userHandler.LoginPage)
	apiv1.POST("/user/login", server.userHandler.LoginByAccount)

	apiv1.Use(server.userHandler.Check())

	apiv1.GET("/user/me", server.userHandler.Me)
	apiv1.POST("/user/admin", server.userHandler.Admin)
	apiv1.GET("/user/logout", server.userHandler.Logout)
	
	apiv1.GET("/ranklist", server.voteHandler.RankList)
	apiv1.GET("/vote", server.voteHandler.GetVoteInfo)
	
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
