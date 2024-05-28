package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"serve/global"
)

//var RabbitMq *amqp.Connection

func InitRabbitmq() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@120.27.159.64:5672/")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	global.RBBITMQ_CON = conn
	fmt.Println("rabbitmq 连接成功")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
