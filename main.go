package main

import (
	"fmt"

	"go-cmdb/app/conf"
	"go-cmdb/app/models"
	router2 "go-cmdb/app/router"
)

func main() {
	// 初始化配置
	conf.InitConf()
	// 初始化数据库连接
	models.InitDB()
	defer models.DB.Close()
	router := router2.InitRouter()
	router.Run(fmt.Sprintf("%s:%s", conf.ServerConf.IP, conf.ServerConf.Port))
}
