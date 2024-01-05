package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"serve/global"
	"serve/model"
	"time"
)

type commonStruct struct {
}

var CommonService commonStruct

func (com *commonStruct) GetToken() (token string, err error) {
	token, err = global.Backend_REDIS.Get(context.Background(), "miniToken").Result()
	fmt.Println("redis--token" + token)
	if true {
		params1 := url.Values{}
		params1.Set("appid", "wxd48baf1e1806673c")
		params1.Set("secret", "3e9992f7101fac422eb7c4629ad72187")
		params1.Set("grant_type", "client_credential")
		apiURL1 := "https://api.weixin.qq.com/cgi-bin/token"

		resp1, err := http.Get(apiURL1 + "?" + params1.Encode())
		if err != nil {
			fmt.Println("请求失败:", err)
			return token, err
		}
		defer resp1.Body.Close()

		// 读取响应内容
		body1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return token, err
		}

		// 打印响应内容
		fmt.Println(string(body1) + "------------")

		tokenInfo := make(map[string]string)
		json.Unmarshal(body1, &tokenInfo)
		fmt.Println(tokenInfo)
		//pubTem(body[""])
		//openid:= body.openid
		token = tokenInfo["access_token"]
		fmt.Println("access_tokencacheset-----------", token)

		// Set the value of the key "foo" to "bar", with the default expiration time
		global.Backend_REDIS.Set(context.Background(), "miniToken", token, 110*time.Minute)
		fmt.Println("response-token" + token)
	}
	return token, err

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

func (com *commonStruct) getHeathData() map[string]SubscribeMessage {
	data := map[string]SubscribeMessage{
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
	}
	return data
}
func (com *commonStruct) getActiveData() map[string]SubscribeMessage {
	data := map[string]SubscribeMessage{
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
	}
	return data
}
func (com *commonStruct) PubClock(clock *model.Clocks) {
	date := clock.TipTime.Format("2006-01-02 15:04:05")
	data := map[string]SubscribeMessage{
		"thing1": SubscribeMessage{
			Value: clock.Title,
		},
		"thing3": SubscribeMessage{
			Value: clock.Describe,
		},
		"time5": SubscribeMessage{
			Value: date,
		},
	}
	com.PubTemCron(data, "XKdx0esytPR0ElXybw-d_0VBBBmP-y8I2w7UV8F9uxk", clock.Openid)
}
func (com *commonStruct) PubTemCron(data map[string]SubscribeMessage, temp_id string, openid string) {

	//openid := c.GetHeader("openid")
	token, err := com.GetToken()

	// 接口 URL
	apiURL := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + token
	// 构建订阅消息
	message := MiniProgramSubscribeMessage{
		ToUser: openid,
		//TemplateID:       "0RWFOTZw9hhlhQh9fLJlmoFoGcuiwpxf3aB3LtQdV2U",
		TemplateID:       temp_id,
		Page:             "pages/cardList/cardList",
		MiniprogramState: "developer",
		Data:             data,
	}

	// 将订阅消息转换为 JSON
	messageData, err := json.Marshal(message)
	fmt.Println(string(messageData))
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
func (com *commonStruct) PubTem(c *gin.Context, data map[string]SubscribeMessage, temp_id string) {

	openid := c.GetHeader("openid")
	token, err := com.GetToken()

	// 接口 URL
	apiURL := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + token
	// 构建订阅消息
	message := MiniProgramSubscribeMessage{
		ToUser: openid,
		//TemplateID:       "0RWFOTZw9hhlhQh9fLJlmoFoGcuiwpxf3aB3LtQdV2U",
		TemplateID:       temp_id,
		Page:             "pages/cardList/cardList",
		MiniprogramState: "developer",
		Data:             data,
	}

	// 将订阅消息转换为 JSON
	messageData, err := json.Marshal(message)
	fmt.Println(string(messageData))
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
