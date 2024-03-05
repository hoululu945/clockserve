package common

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"serve/global"
	"serve/model"
	"strconv"
	"time"
)

type EmailInterfac interface {
	Send(clock *model.Clocks) error
}

func SendMail(mail EmailInterfac) {
	var clock model.Clocks
	global.Backend_DB.Where(" id = ?", 152).First(&clock)
	mail.Send(&clock)

}

type QQMail struct {
}

type WyMail struct {
}

var Wy WyMail

type GoogleMail struct {
}

func (q *QQMail) send(clock *model.Clocks) error {

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
func (q *GoogleMail) send(clock *model.Clocks) error {
	context.WithCancel(context.Background())
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
