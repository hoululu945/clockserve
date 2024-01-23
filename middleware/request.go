package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"serve/global"
	"serve/model"
	"time"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录结束时间
		end := time.Now()

		// 计算请求处理时间
		duration := end.Sub(start)

		// 获取请求方法和参数
		method := c.Request.Method
		params := c.Request.URL.Query()
		var requestLog model.RequestLog
		requestLog.Openid = c.GetHeader("openid")
		requestLog.RequestMethod = method
		marshal, _ := json.Marshal(params)
		requestLog.RequestParam = string(marshal)
		requestLog.StartTime = start
		requestLog.EndTime = end
		requestLog.SubTime = duration
		requestLog.Handler = c.HandlerName()
		requestLog.RequestPath = c.Request.URL.Path
		openid := c.GetHeader("openid")
		var user model.Users
		global.Backend_DB.Find(&user, map[string]interface{}{"mini_openid": openid})
		if user.ID != 0 {
			requestLog.UserId = int(user.ID)
		}
		global.Backend_DB.Create(&requestLog)
		// 打印请求日志
		fmt.Printf("[%s] %s %s %s %s %s\n", end.Format("2006/01/02 - 15:04:05"), c.ClientIP(), method, c.Request.URL.Path, params.Encode(), duration)
	}
}
