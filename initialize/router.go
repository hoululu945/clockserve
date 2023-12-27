package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	//"serve/middleware"
	//"serve/router"
	"github.com/gin-gonic/gin"
	"serve/router"
)

func Routers() *gin.Engine {
	GoodsRouter := router.RouterApp.GoodsRouter
	CardRouter := router.RouterApp.CardRouter
	ClockRouter := router.RouterApp.ClockRouter
	Router := gin.New()
	store := cookie.NewStore([]byte("secret"))
	Router.Use(sessions.Sessions("mysession", store))
	PublicGroup := Router.Group("api")
	GoodsRouter.InitGoodsRouter(PublicGroup)
	UserRouter := router.RouterApp.UserRouter
	UserRouter.InitUserRouter(PublicGroup)
	CardRouter.InitCardRouter(PublicGroup)
	ClockRouter.InitClockRouter(PublicGroup)
	return Router
}
