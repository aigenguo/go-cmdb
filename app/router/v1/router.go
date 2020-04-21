package v1

import (
	"github.com/gin-gonic/gin"
	"go-cmdb/app/apis"
	"go-cmdb/app/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 在main.go使用日志中间件会造成无法写入日志
	router.Use(middleware.LogToFile())
	v1 := router.Group("/v1")
	{
		v1.GET("/assets", apis.Assets)
		v1.POST("/assets", apis.Store)
		v1.PUT("/assets/:id", apis.Update)
		v1.DELETE("/assets/:id", apis.Destroy)
	}
	return router
}
