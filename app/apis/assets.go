package apis

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-cmdb/app/code"
	"go-cmdb/app/models"
)

func Assets(c *gin.Context) {
	var asset models.Asset
	pageStr := c.DefaultQuery("pageNum", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		code.ErrResp(c, 500, "页数转码失败: "+err.Error())
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		code.ErrResp(c, 500, "页面大小转码失败: "+err.Error())
	}
	query := c.Query("query")
	assets, count, err := asset.Assets(int(pageNum), int(pageSize), query)
	if err != nil {
		log.Fatal(err.Error())
		code.ErrResp(c, 500, "查询数据库失败: "+err.Error())
	}
	var H map[string]interface{}
	H = make(map[string]interface{})
	H["assets"] = assets
	H["count"] = count
	code.SuccessResp(c, 200, "", H)
}
