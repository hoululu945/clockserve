package router

type Router struct {
	GoodsRouter
	UserRouter
	CardRouter
	ClockRouter
	SettingRouter
}

var RouterApp = new(Router)
