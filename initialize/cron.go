package initialize

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/smtp"
	"serve/global"
	"serve/model"
	"serve/service/common"
	"strconv"
	"strings"
	"time"
)

func InitCron() {
	//go runScheduledTask()
	//go runCron()
	go subRedisKeyExpir()
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
	c := cron.New()

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
			googleSendMail(&clock)

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
