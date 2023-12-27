package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"serve/global"
	"serve/model"
	"serve/service/common"
	"strconv"
	"time"
)

type Card struct {
	//common2.Qiniu
}

var CardApi Card

func (card *Card) Add(c *gin.Context) {
	var cardModel model.Cards
	c.BindJSON(&cardModel)

	fmt.Println("add", cardModel)
	session := sessions.Default(c)
	session.Set("openid", "sssss2ssss")
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{})
	}
	openid := c.GetHeader("openid")
	cardModel.Openid = openid
	var count int64
	global.Backend_DB.Model(&cardModel).Where("openid=?", openid).Count(&count)
	count++
	mess := "已经打卡" + strconv.Itoa(int(count)) + "次数，再接再厉"
	//longmess := common.Cardinstance.Long(openid, c)
	fmt.Println(count, "tiaoshu")
	if err = global.Backend_DB.Create(&cardModel).Error; err == nil {
		client := global.Backend_REDIS
		today := time.Now()
		t, _ := strconv.Atoi(today.Format("20060102"))
		client.SetBit(c, openid, int64(t), 1)
		longmess := common.Cardinstance.Long(openid, c)
		date := time.Now().Format("2006-01-02 15:04:05")
		data := map[string]common.SubscribeMessage{
			"thing1": common.SubscribeMessage{
				Value: "就要打卡",
			},
			"time2": common.SubscribeMessage{
				Value: date,
			},
			"thing3": common.SubscribeMessage{
				Value: longmess,
			},
			"thing4": common.SubscribeMessage{
				Value: mess,
			},
		}
		common.CommonService.PubTem(c, data, "_-TV_81Or0oM9IqLn2oi5APXx7PS10GrSY8fhS0xeig")
	}

	c.JSON(http.StatusOK, gin.H{"openid": "sssss"})
	return
}

func (card *Card) List(c *gin.Context) {
	openid := c.GetHeader("openid")
	var cardModels []model.Cards
	desc := c.Query("desc")
	fmt.Println(desc)

	global.Backend_DB.Model(&model.Cards{}).Where("`openid`=? and `describe`=?", openid, desc).Order("id desc").Find(&cardModels)

	c.JSON(http.StatusOK, gin.H{"list": cardModels, "headeropenid": openid})
	return
}
func (card *Card) Detail(c *gin.Context) {
	fmt.Println("detail")
	id := c.Query("id")
	var cardModel model.Cards
	global.Backend_DB.First(&cardModel, id)
	c.JSON(http.StatusOK, gin.H{"detail": cardModel})

}
