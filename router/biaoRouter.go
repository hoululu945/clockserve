package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/middleware"
)

type Biao struct {
}

var BiaoRouter Biao

func (b *Biao) InitBiao(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())
	router.POST("add/biao", api.BiaoController.AddNewBiao)

}
