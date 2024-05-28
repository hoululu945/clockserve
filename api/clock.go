package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/smtp"
	"serve/global"
	"serve/model"
	"serve/service/amqp"
	"strconv"
	"strings"
	"time"
)

type Clock struct {
}

var ClockApi Clock

func testMail() {
	// 设置发件人、收件人和邮件内容
	from := "houlu0621@163.com"
	to := []string{"1650221128@qq.com"}
	subject := "示例邮件"
	message := "这是一封示例邮件的正文内容。"

	// 创建邮件内容
	body := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n\r\n" +
		message

	// 连接到SMTP服务器
	smtpHost := "smtp.163.com"
	smtpPort := "587"
	smtpUsername := "houlu0621@163.com"
	smtpPassword := "LLVEGRWZVSQISDOH"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(body))
	if err != nil {
		log.Fatal("邮件发送失败:", err)
	}

	log.Println("邮件发送成功！")
}

func testMail2() {
	// 设置发件人、收件人和邮件内容
	from := "1650221128@qq.com"
	to := []string{"houlu0621@163.com"}
	subject := "示例邮件"
	message := "这是一封示例邮件的正文内容。"

	// 创建邮件内容
	body := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n\r\n" +
		message

	// 连接到SMTP服务器
	smtpHost := "smtp.mailgun.org"
	smtpPort := "587"
	smtpUsername := "1650221128@qq.com"
	smtpPassword := "@#$hoululu945"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(body))
	if err != nil {
		log.Fatal("邮件发送失败:", err)
	}

	log.Println("邮件发送成功！")
}
func test4() {
	// 邮箱账户信息
	email := "houll52120@gmail.com"
	password := "zigseosjiqsvsfpd" // 应用密码，而不是邮箱账户密码

	// 邮件配置
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// 邮件内容
	from := email
	to := "houlu0621@163.com"
	subject := "Test Email2222"
	body := "This is a test email"

	// 邮箱账户信息

	// 邮件配置

	// 邮件内容

	// 邮箱服务器地址
	smtpServer := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	// 邮件凭证
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// 准备邮件内容
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	// 发送邮件
	err := smtp.SendMail(smtpServer, auth, from, []string{to}, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("邮件发送成功！")
}
func test3() {
	// 设置发件人、收件人和邮件内容
	from := "1650221128@qq.com"
	to := []string{"houlu0621@163.com"}
	subject := "示例邮件"
	message := "这是一封示例邮件的正文内容。"

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
		log.Fatal("邮件发送失败:", err)
	}

	log.Println("邮件发送成功！")
}
func (cl Clock) Add(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println("天机-入口")
	loc, err1 := time.LoadLocation("Asia/Shanghai")
	if err1 != nil {
		fmt.Println("无法加载时区:", err1)
		return
	}

	// 获取当前本地时间并转换为北京时间
	now := time.Now().In(loc)
	// 输出结果
	fmt.Println("当前北京时间:", now)
	var Clocks model.Clocks
	var ClocksMap map[string]interface{}
	c.ShouldBindJSON(&ClocksMap)
	tipTimeStr := ClocksMap["tipTime"]
	delete(ClocksMap, "tipTime")
	//localLoc := time.Local
	localTimezone := "Asia/Shanghai"
	loc, err := time.LoadLocation(localTimezone)
	if err != nil {
		fmt.Println("无法加载时区:", err)
		return
	}

	// 解析时间字符串为本地时间
	tipTimeDate, err := time.ParseInLocation("2006-01-02 15:04:05", tipTimeStr.(string), loc)
	fmt.Println(tipTimeDate, "----------------")
	Clocks.TipTime = tipTimeDate
	des := ClocksMap["des"]
	Clocks.Describe = des.(string)
	Clocks.TipImage = ClocksMap["tipImage"].(string)
	openid := c.GetHeader("openid")
	Clocks.Openid = openid
	Clocks.Title = ClocksMap["title"].(string)
	fmt.Println(ClocksMap)
	s := ClocksMap["reminderType"].(string)
	atoi, _ := strconv.Atoi(s)
	Clocks.ReminderType = atoi
	err = global.Backend_DB.Create(&Clocks).Error
	duration := tipTimeDate.Sub(now)
	fmt.Println(tipTimeDate, now)
	fmt.Println(duration)
	err1 = amqp.Publish(Clocks.ID, duration)
	fmt.Println(err1)
	//err = global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(Clocks.ID)), Clocks.ID, duration).Err()
	//fmt.Println("err----", err)

	c.JSON(0, Clocks)
}

func (cl Clock) List(c *gin.Context) {

	var Clocks model.Clocks
	var ClocksArr []model.Clocks
	name := c.Query("name")
	openid := c.GetHeader("openid")

	//remark := c.Query("remark")
	//fmt.Println(name, remark)
	query := global.Backend_DB.Model(&Clocks).Order("id desc").Where("openid=? and type=0", openid).Where("`describe` like ? or `title` like ?", "%"+name+"%", "%"+name+"%")
	query.Limit(50).Find(&ClocksArr)
	fmt.Println(ClocksArr)
	c.JSON(0, ClocksArr)
}
