package core

import (
	"fmt"
	"serve/config"
	"serve/initialize"
	"serve/service/amqp"
)

func RunServer1() {
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
	fmt.Println("************&&&&&&&&&&&66666&&&&&&&&&&&1111111111111111")

	Router.Run(":8082")
	fmt.Println("************&&&&&&&&&&&&&&&&&&&&&&1111111111111111")

}
