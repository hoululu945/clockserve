package router

import (
	"github.com/gin-gonic/gin"
	"serve/api"
	"serve/api/common"
)

type CardRouter struct {
}

func (c *CardRouter) InitCardRouter(group *gin.RouterGroup) {
	group.POST("card/add", api.CardApi.Add)
	group.GET("card/list", api.CardApi.List)
	group.GET("card/detail", api.CardApi.Detail)
	group.GET("qiniu/token", common.QniuApi.GetToken)

}
