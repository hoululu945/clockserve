package common

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ccc(c *gin.Context) {
	date := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]SubscribeMessage{
		"thing1": SubscribeMessage{
			Value: "就要打卡",
		},
		"time2": SubscribeMessage{
			Value: date,
		},
		"thing3": SubscribeMessage{
			Value: "ssss",
		},
		"thing4": SubscribeMessage{
			Value: "mess",
		},
	}
	CommonService.PubTem(c, data, "_-TV_81Or0oM9IqLn2oi5APXx7PS10GrSY8fhS0xeig")
}
