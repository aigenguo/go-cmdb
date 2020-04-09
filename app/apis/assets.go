package apis

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-cmdb/app/models"
)

func Assets(c *gin.Context) {
	var asset models.Asset
	pageStr := c.DefaultQuery("pageNum", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	query := c.Query("query")
	assets, count, err := asset.Assets(int(pageNum), int(pageSize), query)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   1,
		"assets": assets,
		"count":  count,
	})
}
