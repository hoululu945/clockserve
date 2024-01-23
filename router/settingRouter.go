package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/middleware"
)

type SettingRouter struct {
}

func (s *SettingRouter) InitSettingRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())

	router.POST("setting/add", api.SettingApi.Add)
	router.GET("setting/list", api.SettingApi.List)
	router.GET("setting/get", api.SettingApi.Get)
	router.GET("setting/weather", api.SettingApi.Weather)
}
