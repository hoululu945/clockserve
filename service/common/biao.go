package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"serve/global"
	model2 "serve/model"
	"time"
)

type biao struct {
}

var BiaoSerivce biao

type Result struct {
	Categorys   []Category `json:"categorys"`
	Totalcount  int        `json:"totalcount"`
	Records     []Record   `json:"records"`
	ScorllId    string     `json:"scorll_id"`
	Executetime string     `json:"executetime"`
}
type Categorys struct {
	Arrcate []Category `json:"arrcate"`
}
type Category struct {
	Categorynum  string `json:"categorynum"`
	Count        string `json:"count"`
	Categoryname string `json:"categoryname"`
}

//	type Records struct {
//		Record model.Biao `json:"record"`
//	}
type Record struct {
	Categorynum    string `json:"categorynum"`
	Infoid         string `json:"infoid"`
	Sysclicktimes  int    `json:"sysclicktimes"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Webdate        string `json:"webdate"`
	Highlight      `json:"highlight"`
	Projectno      string `json:"projectno"`
	Syscategory    string `json:"syscategory"`
	Syscollectguid string `json:"syscollectguid"`
	Linkurl        string `json:"linkurl"`
	Id             string `json:"id"`
	Sysscore       string `json:"sysscore"`
	Infodate       string `json:"infodate"`
}
type Highlight struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Webdate string `json:"webdate"`
}

type RepStruct struct {
	Result `json:"result"`
}

func (b *biao) clientRequest(keyword string, pn int) RepStruct {
	poststr := `{"token":"","pn":0,"rn":10,"sdt":"","edt":"","wd":"食品shiacaishihhh","inc_wd":"","exc_wd":"","fields":"title;content;projectno","cnum":"","sort":"{\"webdate\":\"0\"}","ssort":"title","cl":500,"terminal":"","condition":[{"fieldName":"categorynum","notEqualList":["002002010"]}],"time":null,"highlights":"title;content","statistics":null,"unionCondition":null,"accuracy":"","noParticiple":"1","searchRange":null}`
	var postMapData map[string]interface{}
	json.Unmarshal([]byte(poststr), &postMapData)
	url := "https://ggzy.hefei.gov.cn/inteligentsearch/rest/esinteligentsearch/getFullTextDataNew"
	//// 发送HTTP GET请求
	client := &http.Client{}
	postMapData["wd"] = keyword
	postMapData["pn"] = pn
	marshal, _ := json.Marshal(postMapData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var rep1 RepStruct

	defer resp.Body.Close()
	if resp.StatusCode == 200 {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应体失败:", err)
		}
		err = json.Unmarshal(body, &rep1)
		if err != nil {
			fmt.Println("解码JSON失败:", err)
		}

	}
	return rep1
}
func (b *biao) AddNewBiao(keyword string) RepStruct {
	//wantCount := 50
	//pagecount := 10
	////num := int(wantCount/10)
	//num := int(math.Ceil(float64(wantCount) / float64(pagecount)))

	wantCount := 10
	pageCount := 10
	count := pageCount
	rep1 := b.clientRequest(keyword, count)
	b.Add(rep1.Records, keyword)

	actualnum := rep1.Totalcount
	if actualnum <= wantCount {
		wantCount = actualnum
	}
	num := (wantCount + pageCount - 1) / pageCount
	time.Sleep(2 * time.Second)
	for i := 1; i < num; i++ {
		count = i*pageCount + pageCount
		fmt.Println("count----------", count)
		rep1 = b.clientRequest(keyword, count)
		b.Add(rep1.Records, keyword)
		time.Sleep(2 * time.Second)

	}
	return rep1

}
func (r *biao) Add(results []Record, types string) {
	var biao model2.Biao
	var biaos []model2.Biao
	for _, v := range results {
		biao.Categorynum = v.Categorynum
		biao.Infoid = v.Infoid
		biao.Syscollectguid = v.Syscollectguid
		biao.Title = v.Title
		biao.Content = v.Content
		parse, _ := time.Parse("2006-01-02 15:04:05", v.Webdate)
		biao.Webdate = parse
		biao.Projectno = v.Projectno
		biao.Syscategory = v.Syscategory
		biao.Syscollectguid = v.Syscollectguid
		biao.Linkurl = v.Linkurl
		biao.BiaoId = v.Id
		biao.Infodate, _ = time.Parse("2006-01-02 15:04:05", v.Infodate)
		biao.Type = types
		b := model2.Biao{}
		res := global.Backend_DB.Where("biao_id=?", v.Id).First(&b)

		if res.RowsAffected == 0 {
			biaos = append(biaos, biao)
		}
	}
	fmt.Println(biaos, "&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	global.Backend_DB.Create(biaos)
}
