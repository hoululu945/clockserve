package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"serve/global"
	"serve/model"
	"strconv"
	"time"
)

type CardService struct {
}

var Cardinstance CardService

func (card CardService) Long(openid string, c *gin.Context) string {
	client := global.Backend_REDIS
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	// 获取当前日期的打卡情况
	t, _ := strconv.Atoi(today.Format("20060102"))
	todayBit := client.GetBit(c, openid, int64(t)).Val()

	// 获取前一天的打卡情况
	y, _ := strconv.Atoi(yesterday.Format("20060102"))
	yesterdayBit := client.GetBit(c, openid, int64(y)).Val()
	var continuousCount int64 = 0
	fmt.Println(int64(t), y, yesterdayBit, todayBit)
	if todayBit == 1 && yesterdayBit == 1 {
		// 如果今天和昨天都打卡了，增加连续打卡天数
		continuousCount = client.Incr(c, openid+"_continuous_count").Val()
		fmt.Println(continuousCount)
	} else {
		if yesterdayBit == 0 {
			if todayBit == 1 {
				client.SetBit(c, openid, int64(t), 1)
				client.Set(c, openid+"_continuous_count", 1, 0)
				continuousCount = 1
			} else {
				client.Set(c, openid+"_continuous_count", 0, 0)

				continuousCount = 0

			}
		}
		// 如果今天或昨天有一天或两天都没有打卡，则将连续打卡天数归零
		client.Set(c, openid+"_continuous_count", 0, 0)
	}
	if continuousCount >= 7 {
		// 将连续打卡天数归零
		client.Set(c, openid+"_continuous_count", 0, 0)

		// 将过去7天的打卡记录归零
		for i := 0; i < 7; i++ {
			resetDay := today.AddDate(0, 0, -i)
			f, _ := strconv.Atoi(resetDay.Format("20060102"))
			client.SetBit(c, openid, int64(f), 0)
		}
	}
	var cardmodel model.Cards
	global.Backend_DB.Where("openid=?", openid).Order("max_card_num desc").First(&cardmodel)
	maxNum := cardmodel.MaxCardNum
	if maxNum < continuousCount {
		cardmodel.MaxCardNum = continuousCount
		maxNum = continuousCount
		global.Backend_DB.Save(&cardmodel)
	}
	sprintf := fmt.Sprintf("连续打卡 %d 天,最多连续打卡次数%d次", continuousCount, maxNum)
	return sprintf
}
