package main

import (
	"fmt"

	"go-cmdb/app/conf"
	orm "go-cmdb/app/database"
	router2 "go-cmdb/app/router"
)

func main() {
	//	基本配置初始化
	conf.InitConf()
	defer orm.DB.Close()
	router := router2.InitRouter()
	router.Run(fmt.Sprintf(":%d", conf.ServerConf.Port))
}
