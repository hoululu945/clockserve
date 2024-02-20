package initialize

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/smtp"
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
func weathertitle(str string) string {
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
	temper, _ := strconv.Atoi(s[index])
	temperLow, _ := strconv.Atoi(s[index-2])

	if temper <= 2 {
		title += "最高温度低于3度记得加衣 "
	}
	fmt.Println(s[index-3], s[index-2], s[index-1], s[index], s[index+1], s, "**********************************")
	if temperLow > 0 && s[index-3] == "-" {
		title += fmt.Sprintf(" 最低温度零下%d摄氏度记得加衣 ", temperLow)
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
			title := weathertitle(weather.Sons[1].Cloud)
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
	go subRedisKeyExpir()
	go everySecond()
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
				sendEmail(&v)

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
		sendEmail(&clo)

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
	subscribe := redis.Subscribe(context.Background(), "__keyevent@0__:expired")
	defer subscribe.Close()

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
			//sendEmail(&clock)
			sendEmail(&clock)

		}
	}
	//subscribe.Close()
	// 关闭订阅者
}

func googleSendMail(clock *model.Clocks) {
	fmt.Println("开始发了google*****************************")
	var user model.Users
	global.Backend_DB.First(&user, map[string]interface{}{"mini_openid": clock.Openid})
	// 邮箱账户信息
	// 邮箱账户信息
	email := "houll52120@gmail.com"
	password := "zigseosjiqsvsfpd" // 应用密码，而不是邮箱账户密码

	// 邮件配置
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// 邮件内容
	from := email
	to := user.Email
	subject := clock.Title
	body := clock.Describe

	// 邮箱账户信息

	// 邮件配置

	// 邮件内容

	// 邮箱服务器地址
	smtpServer := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	// 邮件凭证
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// 准备邮件内容
	//message := []byte(fmt.Sprintf("To: %s\r\n"+
	//	"Subject: %s\r\n"+
	//	"\r\n"+
	//	"%s\r\n", to, subject, body))
	// 准备邮件头部
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = encodeWord(subject) // 编码邮件标题
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=UTF-8"

	var msg strings.Builder
	for k, v := range header {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)
	// 发送邮件
	err := smtp.SendMail(smtpServer, auth, from, []string{to}, []byte(msg.String()))
	if err != nil {
		fmt.Println("邮件发失败！" + err.Error())
	}
	reminderType := clock.ReminderType
	if reminderType != 0 && clock.IsCircle == 1 {
		duration := time.Second
		switch reminderType {
		case 1:
			duration = 24 * 60 * 60 * time.Second
		case 2:
			duration = 24 * 60 * 60 * 7 * time.Second
		case 3:
			duration = 24 * 60 * 60 * 30 * time.Second
		case 4:
			duration = 24 * 60 * 60 * 365 * time.Second
		}
		err = global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(clock.ID)), clock.ID, duration).Err()
		fmt.Println("添加新的循环成功成功！", duration)

	}

	fmt.Println("err----", err)
	clock.IsTip = 1
	global.Backend_DB.Save(clock)
	fmt.Println("邮件发送成功！")
}

// 编码邮件标题为 UTF-8 格式
func encodeWord(word string) string {
	// 使用 base64 进行编码
	encoded := base64.StdEncoding.EncodeToString([]byte(word))
	return "=?UTF-8?B?" + encoded + "?="
}
func sendEmail(clock *model.Clocks) {
	fmt.Println("开始发了*****************************")
	var user model.Users
	global.Backend_DB.First(&user, map[string]interface{}{"mini_openid": clock.Openid})
	// 设置发件人、收件人和邮件内容
	from := "1650221128@qq.com"
	to := []string{user.Email}
	subject := "闹钟提醒"
	message := clock.Describe

	// 创建邮件内容
	body := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n\r\n" +
		message

	// 连接到SMTP服务器
	smtpHost := "smtp.qq.com"
	smtpPort := "587"
	smtpUsername := "1650221128@qq.com"
	smtpPassword := "jpsodytbwthweaga"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(body))
	if err != nil {
		log.Println("邮件发送失败:", err)
	}

	reminderType := clock.ReminderType
	if reminderType != 0 && clock.IsCircle == 1 {
		duration := time.Second
		switch reminderType {
		case 1:
			duration = 24 * 60 * 60 * time.Second
		case 2:
			duration = 24 * 60 * 60 * 7 * time.Second
		case 3:
			duration = 24 * 60 * 60 * 30 * time.Second
		case 4:
			duration = 24 * 60 * 60 * 365 * time.Second
		}
		err = global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(clock.ID)), clock.ID, duration).Err()
		fmt.Println("添加新的循环成功成功！", duration)

	} else {
		clock.IsTip = 1
		global.Backend_DB.Save(clock)
	}

	fmt.Println("err----", err)
	//clock.IsTip = 1
	//global.Backend_DB.Save(clock)

	log.Println("邮件发送成功！")

	// 邮件参数
	//from := "houlu0621@163.com"
	//password := "LLVEGRWZVSQISDOH" //发送方的邮箱密码，注意用163邮箱这里填写的是“客户端授权密码”而不是邮箱的登录密码！
	////to := "1650221128@qq.com"
	//to := user.Email
	//subject := "闹钟提醒"
	//body := clock.Describe
	//
	////SMTP服务器配置
	//smtpHost := "smtp.163.com"
	//smtpPort := "587"
	//
	//// 创建认证
	//auth := smtp.PlainAuth("", from, password, smtpHost)
	//
	//// 组装邮件内容
	//msg := []byte("To: " + to + "\r\n" +
	//	"From: " + from + "\r\n" +
	//	"Subject: " + subject + "\r\n" +
	//	"\r\n" +
	//	body + "\r\n")
	//
	//// 发送邮件
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println("邮件发送成功！")
}
