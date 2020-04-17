package v1

import (
	"github.com/gin-gonic/gin"
	"go-cmdb/app/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/assets", apis.Assets)
		v1.POST("/assets", apis.Store)
		v1.PUT("/assets/:id", apis.Update)
		v1.DELETE("/assets/:id", apis.Destroy)
	}
	return router
}
