package router

import (
	"github.com/gin-gonic/gin"
	"go-cmdb/app/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/assets", apis.Assets)
	router.POST("/assets", apis.Store)
	router.PUT("/assets/:id", apis.Update)
	router.DELETE("/assets/:id", apis.Destroy)
	return router
}
