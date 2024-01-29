package router

type Router struct {
	GoodsRouter
	UserRouter
	CardRouter
	ClockRouter
	SettingRouter
	Biao
}

var RouterApp = new(Router)
