package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"serve/global"
	"time"
)

// GormPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormPgSql() *gorm.DB {
	//p := global.GVA_CONFIG.Pgsql
	//if p.Dbname == "" {
	//	return nil
	//}
	pgsqlConfig := postgres.Config{
		DSN:                  Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志模式为 Info，即打印 SQL 日志)
		NowFunc: func() time.Time {
			return time.Now().In(time.FixedZone("CST", 8*60*60)) // 设置时区为北京时区
		},
	}); err != nil {
		log.Fatal(err)

		return nil
	} else {
		global.Backend_DB = db

		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(11)
		sqlDB.SetMaxOpenConns(11)
		return db
	}
}

func Dsn() string {
	//return "postgres:123098@tcp(120.27.159.64:5432)/inventory?"
	return "host=120.27.159.64  user=postgres  password=123098  dbname=inventory port=5432 sslmode=disable"

}
