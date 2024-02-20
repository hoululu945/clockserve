package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"serve/global"
	"time"
)

func InitMysql() *gorm.DB {

	//dsn := "houguo:10220917@tcp(154.83.15.174:3306)/inventory?charset=utf8&parseTime=True"
	dsn := "root:very_strong_password@tcp(120.27.159.64:3306)/inventory?charset=utf8&parseTime=True"
	//dsn := "root:root@tcp(127.0.0.1:3306)/inventory?charset=utf8&parseTime=True"
	//open := mysql.Open("houguo:10220917@tcp(154.83.15.174:3306)/inventory?charset=utf8&parseTime=True&loc=Local")
	//open := mysql.Open("root:123456@tcp(127.0.0.1:3306)/inventory?charset=utf8&parseTime=True&loc=Local")
	//open := mysql.Open("root:root@tcp(127.0.0.1:3306)/inventory?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志模式为 Info，即打印 SQL 日志)
		NowFunc: func() time.Time {
			return time.Now().In(time.FixedZone("CST", 8*60*60)) // 设置时区为北京时区
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	global.Backend_DB = db
	return db
}
