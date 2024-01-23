package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/middleware"
)

type ClockRouter struct{}

func (u *ClockRouter) InitClockRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())

	router.POST("clock/add", api.ClockApi.Add)
	router.GET("clock/list", api.ClockApi.List)

}
