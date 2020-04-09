package router

import (
	"github.com/gin-gonic/gin"
	"go-cmdb/app/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/assets", apis.Assets)
	return router
}
