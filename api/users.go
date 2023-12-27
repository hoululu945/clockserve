package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"serve/global"
	"serve/model"
)

type UsersStruct struct {
}

var Users1 UsersStruct

func (u *UsersStruct) Add(c *gin.Context) {
	fmt.Println("user-add")
}

type LoginRequest struct {
	Code     string `json:"code"`
	UserInfo struct {
		NickName  string `json:"nickName"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"userInfo"`
}
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (u *UsersStruct) ClockAdd(c *gin.Context) {
	session := sessions.Default(c)
	get := session.Get("openid")
	c.JSON(http.StatusOK, gin.H{"openid": get})
	return

}
func (u *UsersStruct) ClockSave(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("openid", "sssss2ssss")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"openid": "ssss"})
	return

}
func (u *UsersStruct) SetEmail(c *gin.Context) {
	email := c.Query("email")
	openid := c.GetHeader("openid")
	var user model.Users
	user.Email = email
	global.Backend_DB.Where("mini_openid=?", openid).Updates(&user)
	c.JSON(http.StatusOK, "success")
}
func (u *UsersStruct) MiniLogin(c *gin.Context) {

	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		// 请求参数错误处理
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 假设在这里使用临时登录凭证 code 和用户信息进行登录验证
	// 你可以通过调用微信的接口获取用户的 OpenID 和 Session Key
	// 并将其存储在数据库中，用于后续的身份验证和用户信息获取

	// 假设登录验证通过后，返回一个用户标识
	fmt.Println(req)

	// 接口 URL
	apiURL := "https://api.weixin.qq.com/sns/jscode2session"

	// 构建请求参数
	params := url.Values{}
	params.Set("appid", "wxd48baf1e1806673c")
	params.Set("secret", "3e9992f7101fac422eb7c4629ad72187")
	code1 := req.Code
	params.Set("js_code", code1)
	params.Set("grant_type", "authorization_code")

	// 发起 GET 请求
	resp, err := http.Get(apiURL + "?" + params.Encode())
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	var user model.Users
	json.Unmarshal(body, &user)
	user.NickName = req.UserInfo.NickName
	user.AvatarUrl = req.UserInfo.AvatarURL

	where := map[string]string{"mini_openid": user.MiniOpenid}
	err = global.Backend_DB.Order("id desc").First(&user, where).Error
	if err != nil {
		global.Backend_DB.Model(&model.Users{}).Create(&user)
	}
	fmt.Println(user)
	// 获取 Session
	session := sessions.Default(c)
	// 设置 Session 数据
	session.Set("openid", user.MiniOpenid)
	// 保存 Session
	err = session.Save()
	fmt.Println("sessiont-openid---", user.MiniOpenid)
	get := session.Get("openid")
	fmt.Println("getlogin----", get)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	// 打印响应内容
	fmt.Println(string(body) + "******************")
	c.JSON(http.StatusOK, Response{
		0,
		map[string]string{"openid": user.MiniOpenid, "email": user.Email, "nickName": user.NickName},
		"ss",
	})
}
