package api

import (
	"github.com/gin-gonic/gin"
	"serve/service/common"
)

type biao struct {
}

var BiaoController biao

func (b *biao) AddNewBiao(c *gin.Context) {
	//wantCount := 50
	//pagecount := 10
	////num := int(wantCount/10)
	//num := int(math.Ceil(float64(wantCount) / float64(pagecount)))
	//fmt.Println(num) // 输出：3
	//fmt.Println("biao----------------", num)
	//c.JSON(0, gin.H{})
	newBiao := common.BiaoSerivce.AddNewBiao("食品")
	////common.BiaoSerivce.Add(newBiao.Records)
	c.JSON(0, newBiao)
	return
}
