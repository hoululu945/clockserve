package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"serve/global"
	model2 "serve/model"
)

type Setting struct {
}

var SettingApi Setting

func (s *Setting) List(c *gin.Context) {

	fmt.Println("开屏")
	settingModel := model2.Setting{}

	global.Backend_DB.Where("type=1").First(&settingModel)
	c.JSON(http.StatusOK, settingModel)
	return

}
func (s *Setting) Get(c *gin.Context) {

	fmt.Println("开屏")
	settingModel := model2.Setting{}

	global.Backend_DB.Where("type=1").First(&settingModel)
	c.JSON(http.StatusOK, settingModel)
	return

}

func (s *Setting) Add(c *gin.Context) {

	fmt.Println("开屏")
}
