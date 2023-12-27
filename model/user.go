package model

import "serve/global"

type Users struct {
	global.GVA_MODEL
	NickName   string `json:"nickName" gorm:"type:varchar(100);comment:名称"`
	AvatarUrl  string `json:"avatarUrl" gorm:"type:varchar(200);comment:头像"`
	MiniOpenid string `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
	Email      string `json:"email" gorm:"type:varchar(100);comment:邮箱"`
}
