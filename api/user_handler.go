package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"xyzvote/db"
	"xyzvote/types"
)



type UserHandler struct {
	// db
	userStore db.UserStore
	// util
	captcha   *base64Captcha.Captcha
}


func NewUserHandler(store db.UserStore, captcha *base64Captcha.Captcha) *UserHandler {
	return &UserHandler{
		userStore: store,
		captcha:   captcha,
	}
}


func (h *UserHandler) DoLogin(c *gin.Context)  {
	var params types.LoginUserParams
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}
	
	result, err := h.userStore.GetUser(context.Background(), params.Username)
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

	if result.EncryptedPassword != params.Password {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "账号密码错误",
		})
		return
	}

	oneHour := int(time.Second * 60 * 60)
	c.SetCookie("credentials", result.Username, oneHour, "/", "", true, false)  

	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}	


// cookie 校验
func (h *UserHandler) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := c.Cookie("credentials")
		if err != nil || creds == "" {
			c.Redirect(http.StatusPermanentRedirect, "/api/v1/login")
		}

		c.Set("creds", creds)
		c.Next()
	}
}


// jwt 校验
func (h *UserHandler) AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("credentials", "", 3600, "/", "", true, false)
	c.Redirect(http.StatusPermanentRedirect, "/api/v1/login")
}

func (h *UserHandler) Register(c *gin.Context) {
	var params types.RegisterUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	// 密码一致
	if params.Password != params.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "两次密码不同",
		})
		return
	}

	// 密码不能是纯数字
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(params.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码不能为纯数字",
		})
		return
	}

	/*
		由于未设置 unique index 约束，做法倾向于 `先查询，再插入`

     */

	// 校验用户是否存在，这种写法非常不安全。有严重的并发风险
	result1, err := h.userStore.GetUser(context.Background(), params.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 查询失败",
		})
		return
	}

	if result1.Id > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"msg": "用户名已经存在",  
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	var user = types.User{
		UserId:            uuid.New().String(),  // 默认，v4 版本，
		Username:          params.Username,
		EncryptedPassword: string(hashedPassword),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := h.userStore.InsertUser(context.Background(), &user); err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "user.idx_user_id") {
				c.JSON(http.StatusConflict, gin.H{
					"msg": "用户 id 已经存在",
				})
				return
			}
		}

		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "db 操作失败",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"msg": "用户注册成功",
	})
}


func (h *UserHandler) GenerateOTP(c *gin.Context) {
	var (
		id     string
		b64s   string
		answer string
		err    error
	)
	id, b64s, answer, err = h.captcha.Generate()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	log.Println("验证码："+ answer)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"captcha_id": id, 
			"data": b64s,
		},
	})
}


func (h *UserHandler) VerifyOTP(c *gin.Context) {
	var params types.VerifyOTPParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传参错误",
		})
		return
	}

	// 如果设置 clear 为 true，验证无论成功或失败，验证码都作废
	if ok := h.captcha.Verify(params.CaptchaId, params.Data, true); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "验证成功",
	})
}