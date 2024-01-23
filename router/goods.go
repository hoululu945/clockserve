package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"serve/api"
	"serve/middleware"
)

type GoodsRouter struct {
}
type LoginRequest struct {
	Code     string `json:"code"`
	UserInfo struct {
		NickName  string `json:"nickName"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"userInfo"`
}

type LoginResponse struct {
	UserID string `json:"userId"`
}

type tocken struct {
	grant_type string
	appid      string
	secret     string
}

func (u GoodsRouter) InitGoodsRouter(group *gin.RouterGroup) {
	router := group.Use(middleware.RequestLoggerMiddleware())
	router.POST("goods/add", api.Goods.InitGoods)
	router.GET("goods/list", api.Goods.Goods)
	router.GET("goods/test", api.Goods.Test)
	router.POST("goodsImage/upload", api.Goods.UploadFile)
	//router.POST("goodsImage/upload", api.Goods.UploadFile)
	router.POST("login", func(c *gin.Context) {
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
		userID := "example_user_id"

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

		// 打印响应内容
		fmt.Println(string(body) + "******************")

		//get token

		params1 := url.Values{}
		params1.Set("appid", "wxd48baf1e1806673c")
		params1.Set("secret", "3e9992f7101fac422eb7c4629ad72187")
		params1.Set("grant_type", "client_credential")
		apiURL1 := "https://api.weixin.qq.com/cgi-bin/token"

		resp1, err := http.Get(apiURL1 + "?" + params1.Encode())
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		defer resp.Body.Close()

		// 读取响应内容
		body1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return
		}

		// 打印响应内容
		fmt.Println(string(body1) + "------------")

		//openid  token发模板消息
		openInfo := make(map[string]string)
		json.Unmarshal(body, &openInfo)
		fmt.Println(openInfo)
		tokenInfo := make(map[string]string)
		json.Unmarshal(body1, &tokenInfo)
		fmt.Println(tokenInfo)
		//pubTem(body[""])
		//openid:= body.openid
		openid := openInfo["openid"]
		access_token := tokenInfo["access_token"]
		pubTem(openid, access_token)
		// 返回用户标识给小程序
		c.JSON(200, LoginResponse{UserID: userID})
	})

}

type MiniProgramSubscribeMessage struct {
	ToUser           string                      `json:"touser"`
	TemplateID       string                      `json:"template_id"`
	Page             string                      `json:"page"`
	MiniprogramState string                      `json:"miniprogram_state"`
	Data             map[string]SubscribeMessage `json:"data"`
}
type SubscribeMessage struct {
	Value string `json:"value"`
}

func pubTem(openid string, token string) {
	// 接口 URL
	apiURL := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + token

	// 构建订阅消息
	message := MiniProgramSubscribeMessage{
		ToUser:           openid,
		TemplateID:       "0RWFOTZw9hhlhQh9fLJlmoFoGcuiwpxf3aB3LtQdV2U",
		Page:             "pages/index",
		MiniprogramState: "developer",
		Data: map[string]SubscribeMessage{
			"thing1": SubscribeMessage{
				Value: "Hello World",
			},
			"date2": SubscribeMessage{
				Value: "2023-09-08",
			},
			"thing3": SubscribeMessage{
				Value: "Hello World",
			},
			"thing4": SubscribeMessage{
				Value: "Hello World",
			},
		},
	}

	// 将订阅消息转换为 JSON
	messageData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("转换为 JSON 失败:", err)
		return
	}

	// 发起 POST 请求
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(messageData))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)

	// 读取响应内容
	// 这里可以根据实际需求进行处理，例如判断请求是否成功等
	fmt.Println(string(all), "订阅消息发送成功")
}
