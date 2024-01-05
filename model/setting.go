package model

import "serve/global"

type Setting struct {
	global.GVA_MODEL

	Describe string `json:"describe" gorm:"type:varchar(100);comment:描述"`
	ImageUrl string `json:"imageUrl" gorm:"type:varchar(100);comment:图片"`
	Type     int64  `json:"type" gorm:"comment:1开屏"`
	Openid   string `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
}
