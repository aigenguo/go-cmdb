package main

import (
	"fmt"

	"go-cmdb/app/conf"
	orm "go-cmdb/app/database"
	router2 "go-cmdb/app/router"
)

func main() {
	// 初始化配置
	conf.InitConf()
	// 初始化数据库连接
	orm.InitDB()
	defer orm.DB.Close()
	router := router2.InitRouter()
	router.Run(fmt.Sprintf("%s:%s", conf.ServerConf.IP, conf.ServerConf.Port))
}
