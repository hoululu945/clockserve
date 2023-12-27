package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
)

type Qiniu struct {
}

var QniuApi Qiniu

func (q Qiniu) GetToken(c *gin.Context) {
	// 存储相关功能的引入包只有这两个，后面不再赘述

	accessKey := "dPxB3-1_kdUi0uMRLK5rnAy096YEy77V6s2mTYqd"
	secretKey := "11k6MlM6qvGPV7eP4E5IRws0yRksOE0LD86-a9Wt"
	bucket := "min-k"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
	c.JSON(http.StatusOK, gin.H{"token": upToken})

}
