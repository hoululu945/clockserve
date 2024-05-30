package initialize

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"regexp"
	"serve/global"
	"serve/model"
	"serve/service/common"
	"strconv"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai") // 加载本地时区位置
var c = cron.New(cron.WithLocation(loc))

func cronTime() {
	weatherTip()
}
func Weathertitle(str string) string {
	var title string

	s := strings.Split(str, "")
	fmt.Println(s, "title((((")
	index := 0
	for k, v := range s {
		if v == "~" {
			index = k + 1
			break
		}
	}

	//
	//temper, _ := strconv.Atoi(s[index])
	//temperLow, _ := strconv.Atoi(s[index-2])

	// 定义正则表达式，匹配温度数字
	//re := regexp.MustCompile(`(\d+)~(\d+)℃`)
	re := regexp.MustCompile(`(-?\d+)~(-?\d+)℃`)

	// 查找匹配的子串
	match := re.FindStringSubmatch(str)

	// 如果匹配成功，提取温度数字
	if len(match) == 3 {
		temperLow, _ := strconv.Atoi(match[1]) // 最低温度
		temper, _ := strconv.Atoi(match[2])    // 最高温度
		fmt.Printf("最低温度：%s℃，最高温度：%s℃\n", temperLow, temper)
		if temper <= 2 {
			title += "最高温度低于3度记得加衣 "
		}
		fmt.Println(s[index-3], s[index-2], s[index-1], s[index], s[index+1], s, "**********************************")
		if temperLow < 0 {
			title += fmt.Sprintf(" 最低温度零下%d摄氏度记得加衣 ", temperLow)
		}
	} else {
		fmt.Println("未找到温度信息")
	}

	if strings.Contains(str, "雨") {
		title += "有雨出门记得带伞"
		if strings.Contains(str, "中雨") {
			title += "有中雨出门记得带伞"
		}
		if strings.Contains(str, "大雨") {
			title += "有大雨出门记得带伞"
		}
		if strings.Contains(str, "暴雨") {
			title += "有暴雨出门记得带伞"
		}
	}
	if strings.Contains(str, "雪") {
		title += "有雪出门记得带伞"
		if strings.Contains(str, "中雪") {
			title += "有中雪出门记得带伞"
		}
		if strings.Contains(str, "大雪") {
			title += "有大雪出门记得带伞"
		}
		if strings.Contains(str, "暴雪") {
			title += "有暴雪出门记得带伞"
		}
	}

	return title
}

func weatherTip() {
	//loc, err := time.LoadLocation("Asia/Shanghai") // 加载本地时区位置
	//if err != nil {
	//	fmt.Println("加载时区失败:", err)
	//	return
	//}
	//c := cron.New(cron.WithLocation(loc))

	c.AddFunc("0 20 * * *", func() {

		weather := common.WeatherService.Weather("/7/")
		fmt.Println("获取天气长度----", len(weather.Sons))
		if len(weather.Sons) >= 1 {
			var clock model.Clocks
			fmt.Println(weather.Sons[1])
			clock.TipImage = weather.Sons[1].Image
			clock.Openid = ""
			clock.Describe = weather.Sons[1].Date + "" + weather.Sons[1].Cloud
			title := Weathertitle(weather.Sons[1].Cloud)
			fmt.Println("title--------" + title + "*****" + weather.Sons[1].Cloud)
			if title != "" {
				clock.Title = title

				var users []model.Users
				global.Backend_DB.Find(&users)
				for _, v := range users {
					clock.Openid = v.MiniOpenid
					common.WeatherService.Add(clock)
				}
			}
		}

		//fmt.Println("tick every 1 second", weather.Sons[0])

	})

	c.Start()
	time.Sleep(time.Second * 5)
}
func InitCron() {
	//go runScheduledTask()
	//go runCron()
	//定时任务
	go cronTime()
	//go subRedisKeyExpir()
	//go everySecond()
}
func everySecond() {
	//loc, err := time.LoadLocation("Asia/Shanghai") // 加载本地时区位置
	//if err != nil {
	//	fmt.Println("加载时区失败:", err)
	//	return
	//}
	//c := cron.New(cron.WithLocation(loc))
	fmt.Println("begin**************")
	c.AddFunc("@every 1m", func() {
		fmt.Println("begin**************")

		var Clocks []model.Clocks
		//location, _ := time.LoadLocation("Asia/Shanghai")
		//hours := -8 * time.Hour
		//now := time.Now().Add(hours).In(location)
		tipTimeDate := time.Now().Format("2006-01-02 15")

		global.Backend_DB.Where("tip_time=? and is_tip=? and type=?", tipTimeDate+":00:00", 0, 1).Find(&Clocks)
		for _, v := range Clocks {
			if v.ID != 0 {
				SendWangyiMail(&v)

			}
		}

	})

}
func runScheduledTask() {
	for {
		// 定时任务的逻辑
		fmt.Println("Running scheduled task...")
		// 执行你的定时任务操作
		rangeClock()
		// 暂停一段时间，这里使用 time.Sleep() 来模拟定时任务的执行间隔
		time.Sleep(5 * time.Second) // 5分钟
	}
}

func rangeClock() {
	var ClocksArr []model.Clocks
	//remark := c.Query("remark")
	//fmt.Println(name, remark)
	now := time.Now()
	query := global.Backend_DB.Model(&model.Clocks{}).Order("id desc").Where("`tip_time` <= ? and `is_tip` = 0", now)
	query.Find(&ClocksArr)
	fmt.Println(ClocksArr)
	for _, clo := range ClocksArr {
		//clo.I
		//clo.
		clo.IsTip = 1
		global.Backend_DB.Save(clo)
		SendWangyiMail(&clo)

		common.CommonService.PubClock(&clo)
	}
}
func runCron() {
	// 创建一个新的cron调度器
	//c := cron.New()

	// 添加定时任务
	// 每分钟执行一次任务
	c.AddFunc("*/5 * * * *", func() {
		fmt.Println("执行任务")
	})

	// 启动cron调度器
	c.Start()
}
func subRedisKeyExpir() {
	redis := global.Backend_REDIS
	fmt.Println("键过期事件:订阅")

	subscribe := redis.Subscribe(context.Background(), "__keyevent@0__:expired")
	defer subscribe.Close()
	fmt.Println("键过期事件:订阅2")

	// 启动监听
	for msg := range subscribe.Channel() {
		fmt.Println("键过期事件:", msg.Payload)
		stringSplice := strings.Split(msg.Payload, ":")
		fmt.Println(stringSplice)
		switch stringSplice[0] {
		case "clock_id":
			clock := model.Clocks{}
			atoi, _ := strconv.Atoi(stringSplice[1])
			clockId := atoi
			global.Backend_DB.First(&clock, map[string]interface{}{"id": clockId})
			common.CommonService.PubClock(&clock)
			clock.IsTip = 1
			global.Backend_DB.Save(clock)
			//SendWangyiMail(&clock)
			SendWangyiMail(&clock)

		}
	}
	//subscribe.Close()
	// 关闭订阅者
}

func SendWangyiMail(clock *model.Clocks) {
	common.SendMail(new(common.WyMail), clock)
}
