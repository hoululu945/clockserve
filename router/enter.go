package router

type Router struct {
	GoodsRouter
	UserRouter
	CardRouter
	ClockRouter
}

var RouterApp = new(Router)
