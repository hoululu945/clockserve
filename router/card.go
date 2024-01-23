package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/api/common"
	"serve/middleware"
)

type CardRouter struct {
}

func (c *CardRouter) InitCardRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())

	router.POST("card/add", api.CardApi.Add)
	router.GET("card/list", api.CardApi.List)
	router.GET("card/detail", api.CardApi.Detail)
	router.GET("qiniu/token", common.QniuApi.GetToken)

}
