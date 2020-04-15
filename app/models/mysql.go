package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-cmdb/app/conf"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// 连接串
	var conStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBConf.User,
		conf.DBConf.Password,
		conf.DBConf.Host+":"+conf.DBConf.Port,
		conf.DBConf.DBName)
	DB, err = gorm.Open("mysql", conStr)
	if err != nil {
		log.Fatalf("Mysql 连接错误 %v", err)
		// 连接失败，等10秒重新连接
		time.Sleep(10 * time.Second)
		DB, err = gorm.Open("mysql", conStr)
		if err != nil {
			panic(err.Error())
		}
	}

	if DB.Error != nil {
		log.Fatalf("数据库错误 %v", DB.Error)
	}

	// 设置数据库默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DBConf.TablePrefix + defaultTableName
	}

	DB.LogMode(conf.DBConf.Debug)
	// 禁用默认表名的复数形式
	// DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(10)
	// migrate 迁移
	DB.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci",
	).AutoMigrate(&Asset{})
	// 	添加唯一索引
	DB.Model(&Asset{}).AddUniqueIndex("uk_hostname", "hostname")
	DB.Model(&Asset{}).AddUniqueIndex("uk_ip", "ip")
	DB.Model(&Asset{}).AddUniqueIndex("uk_oip", "o_ip")
}
