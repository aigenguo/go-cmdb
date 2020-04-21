package middleware

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-cmdb/app/conf"
)

// 日志记录到文件
func LogToFile() gin.HandlerFunc {
	f, err := os.OpenFile(conf.LogConf.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %s", err.Error())
	}
	// 	实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = f
	gin.DefaultWriter = logger.Out
	// 设置日志级别
	level, err := strconv.Atoi(conf.LogConf.Level)
	if err != nil {
		fmt.Printf("日志格式转码失败: %s", err.Error())
	}
	logger.SetLevel(logrus.Level(level))
	// 	设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// // json方式的请求body
		// body, err := ioutil.ReadAll(c.Request.Body)
		// if err != nil {
		// 	fmt.Printf("请求body读取失败: %s", body)
		// }
		// form方式的请求body
		c.Request.ParseForm()
		body := c.Request.PostForm
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"req_body":     body,
		}).Info()
	}
}
