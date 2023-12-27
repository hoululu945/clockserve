package core

import "serve/initialize"

func RunServer1() {
	Router := initialize.Routers()
	initialize.InitRedis1()
	initialize.InitCron()

	Router.Run(":8082")
}
