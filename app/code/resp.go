package code

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

// ErrResp 错误返回值
func ErrResp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp(ctx, code, msg, data...)
}

// SuccessResp 正确返回值
func SuccessResp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp(ctx, 0, msg, data...)
}

// resp 返回
func resp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp := Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	if len(data) == 1 {
		resp.Data = data[0]
	}
	ctx.JSON(http.StatusOK, resp)
}
