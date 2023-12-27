package initialize

import (
	//"awesomeProject/model"
	"gorm.io/gorm"
	"serve/model"
)

func RegisterTables(db *gorm.DB) {

	db.AutoMigrate(
		// 系统模块表
		model.Goods{},
		model.Users{},
		model.Cards{},
		model.Clocks{},
		//model.Language{},
		//model.Authority{},
		//model.Menu{},
		//model.Meta{},
		//model.MenuBtn{},
		//model.SysBaseMenuParameter{},
	)

}
