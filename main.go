package main

import (
	"gorm.io/gorm"
	"serve/core"
	"serve/initialize"
)

var conn *gorm.DB

func main() {
	conn = initialize.InitMysql()
	if conn != nil {
		//conn.AutoMigrate(model.User{})
		db, _ := conn.DB()
		defer db.Close()
		initialize.RegisterTables(conn)
	}
	//initialize.Routers()
	core.RunServer1()
}
