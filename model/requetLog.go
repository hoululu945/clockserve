package model

import (
	"serve/global"
	"time"
)

type RequestLog struct {
	global.GVA_MODEL
	RequestMethod string        `json:"request_method" gorm:"type:varchar(100);comment:请求方法"`
	RequestParam  string        `json:"request_param" gorm:"type:varchar(200);comment:请求参数"`
	StartTime     time.Time     `json:"start_time" gorm:"comment:开始时间"`
	EndTime       time.Time     `json:"end_time" gorm:"comment:结束时间"`
	UserId        int           `json:"user_id" gorm:"comment:用户id"`
	Openid        string        `json:"openid" gorm:"type:varchar(100);comment:小程序openid"`
	Handler       string        `json:"handler" gorm:"type:varchar(100);comment:请求函数"`
	RequestPath   string        `json:"request_path" gorm:"type:varchar(100);comment:请求路径"`
	SubTime       time.Duration `json:"sub_time" gorm:"comment:请求时间"`
}
