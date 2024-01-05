package main

import (
	"fmt"
	"gorm.io/gorm"
	"serve/core"
	"serve/initialize"
)

var conn *gorm.DB

type test struct {
	Name string
}

var tt *test

func main() {
	fmt.Println(tt)
	return
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
