package model

import (
	"serve/global"
	"time"
)

type Clocks struct {
	global.GVA_MODEL
	Title        string    `json:"title" gorm:"type:varchar(100);comment:标题;" binding:"required"`
	Describe     string    `json:"des" gorm:"type:varchar(100);comment:描述;" binding:"required"`
	TipTime      time.Time `json:"tipTime" gorm:"comment:闹钟时间" binding:"required"`
	Openid       string    `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
	TipImage     string    `json:"tipImage" gorm:"type:varchar(100);comment:闹钟图片"`
	IsTip        int       `json:"isTip" gorm:"comment:是否提示1提示;default:0"`
	Type         int       `json:"type" gorm:"comment:1天气;default:0"`
	IsCircle     int       `json:"isCircle" gorm:"comment:是否循环1循环开始0结束;default:1"`
	ReminderType int       `json:"reminderType" gorm:"comment:0定时1每天2每周3每月4每年;default:0"`
}
