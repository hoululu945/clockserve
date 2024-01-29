package model

import (
	"time"
)
import "serve/global"

type Biao struct {
	global.GVA_MODEL
	Title          string    `json:"title" gorm:"type:text;comment:标题;" binding:"required"`
	Content        string    `json:"content" gorm:"type:text;comment:内容;" binding:"required"`
	BiaoId         string    `json:"biaoId" gorm:"comment:是否提示1提示;default:0"`
	Sysclicktimes  int       `json:"sysclicktimes" gorm:"comment:是否提示1提示;default:0"`
	Categorynum    string    `json:"categorynum" gorm:"type:varchar(100)"`
	Infoid         string    `json:"infoid" gorm:"type:varchar(100)"`
	Webdate        time.Time `json:"webdate" gorm:"type:datetime" binding:"required"`
	Infodate       time.Time `json:"infodate" gorm:"type:datetime" binding:"required"`
	Syscategory    string    `json:"syscategory" gorm:"type:varchar(100)"`
	Syscollectguid string    `json:"syscollectguid" gorm:"type:varchar(100)"`
	Linkurl        string    `json:"linkurl" gorm:"type:varchar(100);comment:链接"`
	Projectno      string    `json:"projectno" gorm:"type:varchar(100);comment:"`
	Type           string    `json:"type" gorm:"type:varchar(100);comment:食品、食材、食堂"`
	IsTip          int       `json:"isTip" gorm:"comment:是否提示1提示;default:0"`
}
