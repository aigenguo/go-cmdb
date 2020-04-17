package main

import (
	"fmt"

	"go-cmdb/app/conf"
	"go-cmdb/app/models"
	"go-cmdb/app/router/v1"
)

func main() {
	// 初始化配置
	conf.InitConf()
	// 初始化数据库连接
	models.InitDB()
	defer models.DB.Close()
	router := v1.InitRouter()
	router.Run(fmt.Sprintf("%s:%s", conf.ServerConf.IP, conf.ServerConf.Port))
}
