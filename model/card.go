package model

import "serve/global"

type Cards struct {
	global.GVA_MODEL
	Lat        string `json:"lat" gorm:"type:varchar(100);comment:经度"`
	Lon        string `json:"lon" gorm:"type:varchar(100);comment:纬度"`
	Describe   string `json:"describe" gorm:"type:varchar(100);comment:描述"`
	Openid     string `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
	CardImage  string `json:"cardImage" gorm:"type:varchar(100);comment:打卡图片"`
	MaxCardNum int64  `json:"maxCardNum" gorm:"comment:最多打卡次数"`
}
