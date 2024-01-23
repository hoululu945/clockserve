package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/middleware"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())

	router.POST("user/add", api.Users1.Add)
	router.POST("user/miniLogin", api.Users1.MiniLogin)
	router.GET("user/clock", api.Users1.ClockAdd)
	router.GET("user/clockSet", api.Users1.ClockSave)
	router.GET("user/setEmail", api.Users1.SetEmail)
	router.GET("user/detail", api.Users1.Detail)
	router.POST("user/update", api.Users1.Update)

}
