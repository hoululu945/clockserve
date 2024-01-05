package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
)

type SettingRouter struct {
}

func (s *SettingRouter) InitSettingRouter(group *gin.RouterGroup) {
	group.POST("setting/add", api.SettingApi.Add)
	group.GET("setting/list", api.SettingApi.List)
	group.GET("setting/get", api.SettingApi.Get)
}
