package conf

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

// server 服务基本配置
type server struct {
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
	RunMode string `mapstructure:"runMode"`
	Static  string `mapstructure:"static"`
}

// server 实例化
var ServerConf = &server{}

// 数据库 配置
type database struct {
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	DBName      string `mapstructure:"dbName"`
	TablePrefix string `mapstructure:"tablePrefix"`
	Debug       bool   `mapstructure:"debug"`
}

// database 实例化
var DBConf = &database{}

// 日志 配置
type logger struct {
	Level  string `mapstructure: "level"`
	Output string `mapstructure: "output"`
}

// logger 实例化
var LogConf = &logger{}

// 生成服务配置
func InitConf() {
	viper.SetConfigType("YAML")
	// 	读取配置文件
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("读取 'config.yaml' 失败: %v\n", err)
	}
	// 	配置内容解析
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("解析 'config.yaml' 失败: %v\n", err)
	}
	// 解析配置赋值
	err = viper.UnmarshalKey("server", ServerConf)
	if err != nil {
		log.Fatalf("解析 'server配置' 失败: %v\n", err)
	}
	err = viper.UnmarshalKey("database", DBConf)
	if err != nil {
		log.Fatalf("解析 'database' 失败: %v\n", err)
	}
	err = viper.UnmarshalKey("log", LogConf)
	if err != nil {
		log.Fatalf("解析 'log' 失败: %v\n", err)
	}
}
