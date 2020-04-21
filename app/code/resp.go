package code

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var result interface{}
	if len(data) == 1 {
		result = data[0]
	} else {
		result = data
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": result,
	})
}
