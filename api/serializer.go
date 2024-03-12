package api


import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
)


/*
	序列化，函数式编程设计风格，option 模式

	{
		"code": 10001,
		"msg":  "success",
		"data": {}
	}

 */


type JSONVal struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data any     `json:"data,omitempty"`
}

type JsonOption func(*JSONVal)

func WithCode(code int) JsonOption {
	return func(v *JSONVal) {
		v.Code = code
	}
}

func WithMsg(msg string) JsonOption {
	return func(v *JSONVal) {
		v.Msg = msg 
	}
}

func WithData(data any) JsonOption {
	return func(v *JSONVal) {
		v.Data = data
	}
}

func JSONOk(ctx *gin.Context, opts ...JsonOption) {
	v := JSONVal{
		Code: 0,
		Msg:  "success",
	}

	for _, opt := range opts {
		opt(&v)
	}

	ctx.JSON(http.StatusOK, v)
}

func JSONErr(ctx *gin.Context, opts ...JsonOption) {
	v := JSONVal{
		Code: 1,
		Msg:  "error",
	}

	for _, opt := range opts {
		opt(&v)
	}

	ctx.JSON(http.StatusOK, v)
}

func JSONErrCustom(ctx *gin.Context, code int, opts ...JsonOption) {
	v := JSONVal{
		Code: 1,
		Msg:  "error",
	}

	for _, opt := range opts {
		opt(&v)
	}

	ctx.JSON(code, v)
}


type User struct {
    Id                 int64     `json:"id,omitempty" gorm:"column:id;primary_key"`
    UserId             string    `json:"user_id"      gorm:"column:user_id"`
    Username           string    `json:"username"     gorm:"column:username"`
    EncryptedPassword  string    `json:"-"            gorm:"column:password"`
    Email              string    `json:"email"        gorm:"column:email"`
    Gender             int8      `json:"gender"       gorm:"column:gender"`
    CreatedAt          time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt          time.Time `json:"updated_at"   gorm:"column:updated_at"`
}
