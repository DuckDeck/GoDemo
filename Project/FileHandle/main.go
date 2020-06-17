package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// "path"
	// "path/filepath"
	// "strings"
	// "github.com/PuerkitoBio/goquery"
	// _ "github.com/go-sql-driver/mysql"
	// "golang.org/x/text/encoding/simplifiedchinese"
)

var (
	uploadFileKey = "upload-key"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/", http.Dir("./img"))
	r.POST("/upload", uploadHandler)
	r.Run("0.0.0.0:9090")
}

func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code": -11,
			"msg":  "request not contail file",
		})
		return
	}

	dst := strconv.FormatInt(time.Now().Unix(), 10) + header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, "img/"+dst); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code": -10,
			"msg":  "save file fail",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "http://lovelive.ink:10087/" + dst,
		})
	}
}
