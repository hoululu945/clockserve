package common

import (
	"context"
	"encoding/base64"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net/smtp"
	"serve/global"
	"serve/model"
	"strconv"
	"strings"
	"time"
)

// 封装
type EmailInterfac interface {
	Send(clock *model.Clocks) error
}

func SendMail(mail EmailInterfac, clock *model.Clocks) {
	mail.Send(clock)
}

type QQMail struct {
}

type WyMail struct {
}

var Wy WyMail

type GoogleMail struct {
}

func (q *QQMail) Send(clock *model.Clocks) error {
	fmt.Println("开始发了*****************************")
	var user model.Users
	global.Backend_DB.First(&user, map[string]interface{}{"mini_openid": clock.Openid})
	// 设置发件人、收件人和邮件内容
	from := "1650221128@qq.com"
	to := []string{user.Email}
	subject := clock.Title
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
	CircleSet(clock)

	log.Println("邮件发送成功！")

	return nil
}
func (q *WyMail) Send(clock *model.Clocks) error {
	fmt.Println("开始发了wy*****************************")
	var user model.Users
	global.Backend_DB.First(&user, map[string]interface{}{"mini_openid": clock.Openid})

	msg := gomail.NewMessage()
	msg.SetHeader("From", "houlu0621@163.com")
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", clock.Title)
	msg.SetBody("text/html", clock.Describe)
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.163.com", 465, "houlu0621@163.com", "LLVEGRWZVSQISDOH")
	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		log.Println("邮件发送失败:", err)
	}

	CircleSet(clock)
	log.Println("邮件发送成功！")
	return nil
}
func encodeWord(word string) string {
	// 使用 base64 进行编码
	encoded := base64.StdEncoding.EncodeToString([]byte(word))
	return "=?UTF-8?B?" + encoded + "?="
}
func (q *GoogleMail) Send(clock *model.Clocks) error {
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
	CircleSet(clock)
	fmt.Println("邮件发送成功！")
	return nil
}

func CircleSet(clock *model.Clocks) {
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
		err := global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(clock.ID)), clock.ID, duration).Err()
		if err == nil {
			fmt.Println("wy添加新的循环成功成功！", duration)
		}

	}
	clock.IsTip = 1
	global.Backend_DB.Save(clock)

}
