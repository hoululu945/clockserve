package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Rabbitmq string
)

func InitRabbitmq() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		panic(err)
	}
	LoadRabbitmqConfig(file)

}

func LoadRabbitmqConfig(file *ini.File) {
	Rabbitmq = file.Section("rabbitmq").Key("Rabbitmq").String()
	fmt.Println("连接协议----" + Rabbitmq)
}
