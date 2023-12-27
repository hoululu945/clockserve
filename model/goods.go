package model

import "serve/global"

type Goods struct {
	global.GVA_MODEL
	Name   string `json:"name" gorm:"type:varchar(100);comment:名称"`
	Remark string `json:"remarks" gorm:"type:varchar(100);comment:备注"`
	Url    string `json:"imageUrl" gorm:"type:varchar(100);comment:图片地址"`
	Num    int    `json:"num" gorm:"comment:数量"`
	Openid string `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
}
