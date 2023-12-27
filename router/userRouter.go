package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(group *gin.RouterGroup) {
	group.POST("user/add", api.Users1.Add)
	group.POST("user/miniLogin", api.Users1.MiniLogin)
	group.GET("user/clock", api.Users1.ClockAdd)
	group.GET("user/clockSet", api.Users1.ClockSave)
	group.GET("user/setEmail", api.Users1.SetEmail)

}
