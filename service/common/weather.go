package common

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mozillazg/go-pinyin"
	"log"
	"net/http"
	"serve/global"
	model2 "serve/model"
	"strconv"
	"strings"
	"time"
)

type weatherStruct struct {
}

var WeatherService weatherStruct

type sonWeather struct {
	Title string
	Date  string
	Cloud string
	Image string
}
type weather struct {
	Title        string
	TodayWeather string
	TodayDate    string
	Sons         []sonWeather
}

func (w *weatherStruct) Weather(days string) weather {
	city := "瑶海"
	fmt.Println("city", city)

	var weatherInfo weather
	var settingModel model2.Setting
	sons := make([]sonWeather, 0)
	url := "https://www.tianqi.com/"
	//city := "hefei"
	//days := "/40/"
	//openid := c.GetHeader("openid")
	//settingModel.Openid = openid
	global.Backend_DB.Order("id asc").Find(&settingModel)
	if city == "" {
		city = settingModel.Area
	}
	fmt.Println(city)
	convert := pinyin.Convert(city, nil)
	var strSlice []string
	for _, subSlice := range convert {
		strSlice = append(strSlice, strings.Join(subSlice, ""))
	}
	cityPinyin := strings.Join(strSlice, "")
	// 发送HTTP GET请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+cityPinyin+days, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		// 使用goquery解析网页内容
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// 获取天气信息
		weatherInfo.Title = strings.ReplaceAll(strings.ReplaceAll(doc.Find("div.weaone_b h1").Text(), " ", ""), "\n", "")
		weatherInfo.TodayWeather = strings.ReplaceAll(strings.ReplaceAll(doc.Find("div.weaone_ba").Text(), " ", ""), "\n", "")
		weatherInfo.TodayDate = strings.ReplaceAll(strings.ReplaceAll(doc.Find("div.weaone_bb").Text(), " ", ""), "\n", "")
		//allwea := strings.TrimSpace(doc.Find("ul.weaul li").Text())
		doc.Find("ul.weaul li").Each(func(i int, ss *goquery.Selection) {
			var son sonWeather
			title := ss.Find("a").AttrOr("title", "")
			son.Title = strings.ReplaceAll(strings.ReplaceAll(title, " ", ""), "\n", "")
			date := ss.Find("a div.weaul_q").Text()
			son.Date = strings.ReplaceAll(strings.ReplaceAll(date, " ", ""), "\n", "")

			cloud := ss.Find("a div.weaul_z").Text()
			son.Cloud = strings.ReplaceAll(strings.ReplaceAll(cloud, " ", ""), "\n", "")
			son.Image = w.weatherImage(son.Cloud)
			sons = append(sons, son)
		})
		weatherInfo.Sons = sons

	} else {
		fmt.Println("Failed to retrieve weather information.")
	}

	return weatherInfo
}

func (w *weatherStruct) weatherImage(str string) string {
	var imageHash string
	if strings.Contains(str, "多云") {
		imageHash = "微信图片_20240116145431.png"
	}
	if strings.Contains(str, "晴") {
		imageHash = "微信图片_20240116145508.png"
	}
	if strings.Contains(str, "雪") {
		imageHash = "微信图片_20240116145515.png"
	}
	if strings.Contains(str, "阴") {
		imageHash = "微信图片_20240116145422.png"
	}
	if strings.Contains(str, "雨") {
		imageHash = "微信图片_20240116145447.png"
	}
	return "http://s687dm7qx.hn-bkt.clouddn.com/" + imageHash
}

func (w *weatherStruct) Add(clockData model2.Clocks) {
	loc, err1 := time.LoadLocation("Asia/Shanghai")
	if err1 != nil {
		fmt.Println("无法加载时区:", err1)
		return
	}

	// 获取当前本地时间并转换为北京时间
	now := time.Now().In(loc)
	// 输出结果
	fmt.Println("当前北京时间:", now)
	var Clocks1 model2.Clocks
	var Clocks2 model2.Clocks

	tipTimeStr1 := time.Now().In(loc).Format("2006-01-02") + " 22:00:00"
	tipTimeStr2 := time.Now().In(loc).Add(24*time.Hour).Format("2006-01-02") + " 07:00:00"
	fmt.Println(tipTimeStr1, tipTimeStr2)
	localTimezone := "Asia/Shanghai"
	loc, err := time.LoadLocation(localTimezone)
	if err != nil {
		fmt.Println("无法加载时区:", err)
		return
	}

	// 解析时间字符串为本地时间
	tipTimeDate1, err := time.ParseInLocation("2006-01-02 15:04:05", tipTimeStr1, loc)
	tipTimeDate2, err := time.ParseInLocation("2006-01-02 15:04:05", tipTimeStr2, loc)

	Clocks1.TipTime = tipTimeDate1
	Clocks1.Describe = clockData.Describe
	Clocks1.TipImage = clockData.TipImage
	Clocks1.Openid = clockData.Openid
	Clocks1.Title = "明天" + clockData.Title
	Clocks1.ReminderType = 0
	Clocks1.Type = 1
	err = global.Backend_DB.Create(&Clocks1).Error
	duration := tipTimeDate1.Sub(now)
	fmt.Println(duration)
	err = global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(Clocks1.ID)), Clocks1.ID, duration).Err()
	//第二天提醒
	fmt.Println(Clocks1)
	Clocks2 = Clocks1
	Clocks2.TipTime = tipTimeDate2
	Clocks2.ReminderType = 0
	Clocks2.ID = 0
	Clocks2.Title = "今天" + clockData.Title
	Clocks2.Type = 1
	err = global.Backend_DB.Create(&Clocks2).Error
	fmt.Println(Clocks2)
	duration2 := tipTimeDate2.Sub(now)
	fmt.Println(duration2)
	err = global.Backend_REDIS.Set(context.Background(), "clock_id:"+strconv.Itoa(int(Clocks2.ID)), Clocks1.ID, duration).Err()

}
