package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
)

type ClockRouter struct{}

func (u *ClockRouter) InitClockRouter(group *gin.RouterGroup) {
	group.POST("clock/add", api.ClockApi.Add)
	group.GET("clock/list", api.ClockApi.List)

}
