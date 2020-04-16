package apis

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-cmdb/app/code"
	"go-cmdb/app/models"
)

// 添加资产
func Store(c *gin.Context) {
	var asset models.Asset
	// models.Asset 不支持取索引
	// for k, v := range c.Request.PostForm {
	// }
	asset.IP = c.Request.FormValue("ip")
	asset.OIP = c.Request.FormValue("oip")
	asset.Hostname = c.Request.FormValue("hostname")
	asset.Host = c.Request.FormValue("host")
	cpu, err := strconv.Atoi(c.Request.FormValue("cpu"))
	if err != nil {
		code.ErrResp(c, 500, "CPU核数转码失败: "+err.Error())
		return
	}
	asset.CPU = int8(cpu)
	mem, err := strconv.Atoi(c.Request.FormValue("mem"))
	if err != nil {
		code.ErrResp(c, 500, "内存转码失败: "+err.Error())
		return
	}
	asset.Mem = int16(mem)
	asset.Disk = c.Request.FormValue("disk")
	bandwidth, err := strconv.Atoi(c.Request.FormValue("bandwidth"))
	if err != nil {
		code.ErrResp(c, 500, "带宽转码失败: "+err.Error())
		return
	}
	asset.Bandwidth = int16(bandwidth)
	asset.OS = c.Request.FormValue("os")
	asset.Principal = c.Request.FormValue("principal")
	asset.UseOf = c.Request.FormValue("use_of")

	id, err := asset.Insert()
	if err != nil {
		code.ErrResp(c, 500, "插入失败: "+err.Error())
		return
	}
	code.SuccessResp(c, 200, "插入成功", id)
	return
}

// 删除资产
func Destroy(c *gin.Context) {
	var asset models.Asset
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := asset.Destroy(id)
	if err != nil || result.ID == 0 {
		code.ErrResp(c, 500, "删除资产失败: "+err.Error())
		return
	}
	code.SuccessResp(c, 200, "删除成功", id)
	return
}

// 修改资产
func Update(c *gin.Context) {
	var asset models.Asset
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	asset.IP = c.Request.FormValue("ip")
	asset.OIP = c.Request.FormValue("oip")
	asset.Hostname = c.Request.FormValue("hostname")
	asset.Host = c.Request.FormValue("host")
	cpu, err := strconv.Atoi(c.Request.FormValue("cpu"))
	if err != nil {
		code.ErrResp(c, 500, "CPU核数转码失败: "+err.Error())
		return
	}
	asset.CPU = int8(cpu)
	mem, err := strconv.Atoi(c.Request.FormValue("mem"))
	if err != nil {
		code.ErrResp(c, 500, "内存转码失败: "+err.Error())
		return
	}
	asset.Mem = int16(mem)
	asset.Disk = c.Request.FormValue("disk")
	bandwidth, err := strconv.Atoi(c.Request.FormValue("bandwidth"))
	if err != nil {
		code.ErrResp(c, 500, "带宽转码失败: "+err.Error())
		return
	}
	asset.Bandwidth = int16(bandwidth)
	asset.OS = c.Request.FormValue("os")
	asset.Principal = c.Request.FormValue("principal")
	asset.UseOf = c.Request.FormValue("use_of")

	result, err := asset.Update(id)
	if err != nil || result.ID == 0 {
		code.ErrResp(c, 500, "修改资产失败: "+err.Error())
		return
	}
	code.SuccessResp(c, 200, "修改成功", id)
	return
}

// 获取资产
func Assets(c *gin.Context) {
	var asset models.Asset
	pageStr := c.DefaultQuery("pageNum", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		code.ErrResp(c, 500, "页数转码失败: "+err.Error())
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		code.ErrResp(c, 500, "页面大小转码失败: "+err.Error())
		return
	}
	query := c.Query("query")
	assets, count, err := asset.Assets(int(pageNum), int(pageSize), query)
	if err != nil {
		log.Fatal(err.Error())
		code.ErrResp(c, 500, "查询数据库失败: "+err.Error())
		return
	}
	var H map[string]interface{}
	H = make(map[string]interface{})
	H["assets"] = assets
	H["count"] = count
	code.SuccessResp(c, 200, "查询成功", H)
	return
}
