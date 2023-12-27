package model

import "serve/global"

type Tips struct {
	global.GVA_MODEL
	Name   string `json:"name" gorm:"type:varchar(100);comment:名称"`
	Remark string `json:"remarks" gorm:"type:varchar(100);comment:备注"`
	Url    string `json:"imageUrl" gorm:"type:varchar(100);comment:图片地址"`
	Num    int    `json:"num" gorm:"comment:数量"`
	IsTip  int    `json:"isTip" gorm:"comment:是否提示1提示;default:0"`
}
