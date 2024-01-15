package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"serve/global"
	model2 "serve/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Setting struct {
}

var SettingApi Setting

func (s *Setting) List(c *gin.Context) {

	fmt.Println("开屏")
	settingModel := model2.Setting{}

	global.Backend_DB.Where("type=1").First(&settingModel)
	c.JSON(http.StatusOK, settingModel)
	return

}
func (s *Setting) Get(c *gin.Context) {

	fmt.Println("开屏")
	settingModel := model2.Setting{}

	global.Backend_DB.Where("type=1").First(&settingModel)
	c.JSON(http.StatusOK, settingModel)
	return

}

func (s *Setting) Add(c *gin.Context) {

	fmt.Println("开屏")
}

type sonWeather struct {
	Title string
	Date  string
	Cloud string
}
type weather struct {
	Title        string
	TodayWeather string
	TodayDate    string
	Sons         []sonWeather
}

func (s *Setting) Weather(c *gin.Context) {
	var weatherInfo weather
	sons := make([]sonWeather, 0)
	url := "https://www.tianqi.com/"
	city := "hefei"
	days := "/7/"

	// 发送HTTP GET请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+city+days, nil)
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
		doc.Find("ul.weaul li").Each(func(i int, s *goquery.Selection) {
			var son sonWeather
			title := s.Find("a").AttrOr("title", "")
			son.Title = strings.ReplaceAll(strings.ReplaceAll(title, " ", ""), "\n", "")
			date := s.Find("a div.weaul_q").Text()
			son.Date = strings.ReplaceAll(strings.ReplaceAll(date, " ", ""), "\n", "")

			cloud := s.Find("a div.weaul_z").Text()
			son.Cloud = strings.ReplaceAll(strings.ReplaceAll(cloud, " ", ""), "\n", "")
			sons = append(sons, son)
			// Extract the text or attribute values from each <li> element
			//text := s.Text()
			//href, _ := s.Attr("href")
			//
			//// Print the extracted values
			//fmt.Printf("Text: %s, Href: %s\n", text, href)
		})
		weatherInfo.Sons = sons

	} else {
		fmt.Println("Failed to retrieve weather information.")
	}
	c.JSON(http.StatusOK, weatherInfo)
	return
}
