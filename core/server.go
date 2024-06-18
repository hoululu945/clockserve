package core

import (
	"fmt"
	"serve/config"
	"serve/initialize"
	"serve/pkg/logger"
	"serve/service/amqp"
)

func RunServer1() {
	logger.Init()
	Router := initialize.Routers()
	initialize.InitRedis1()
	initialize.InitCron()
	config.InitRabbitmq()
	initialize.InitRabbitmq()
	go func() {
		err := amqp.NewConsumer()
		if err != nil {
			fmt.Errorf("new consumer error--------------%s", err)
		}
	}()

	Router.Run(":8082")

}
