package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"serve/global"
	"serve/model"
)

type GoodsStruct struct {
}

var Goods GoodsStruct

func (u *GoodsStruct) Test(c *gin.Context) {
	session := sessions.Default(c)
	get := session.Get("openid")
	c.JSON(http.StatusOK, gin.H{"openid": get})
}
func (u *GoodsStruct) InitGoods(c *gin.Context) {
	var Goods model.Goods
	c.ShouldBindJSON(&Goods)
	global.Backend_DB.Create(&Goods)

	c.JSON(0, Goods)

}
func (u *GoodsStruct) Goods(c *gin.Context) {

	var Goods model.Goods
	var GoodsArr []model.Goods
	name := c.Query("name")
	//remark := c.Query("remark")
	//fmt.Println(name, remark)
	openid := c.GetHeader("openid")

	query := global.Backend_DB.Model(&Goods).Order("id desc").Where("openid=?", openid).Where("name like ?  or remark like ?", "%"+name+"%", "%"+name+"%")
	query.Limit(50).Find(&GoodsArr)
	fmt.Println(GoodsArr)
	c.JSON(0, GoodsArr)
}

func (u *GoodsStruct) UploadFile(c *gin.Context) {
	fmt.Println("开始上传")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存文件到服务器
	err = c.SaveUploadedFile(file, "/var/www/uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功", "imageUrl": "http://120.27.159.64/uploads/" + file.Filename})
}
