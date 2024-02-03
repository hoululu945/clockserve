package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/middleware"
)

type ClockRouter struct{}

func (u *ClockRouter) InitClockRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Openid")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	router.POST("clock/add", api.ClockApi.Add)
	router.GET("clock/list", api.ClockApi.List)

}
