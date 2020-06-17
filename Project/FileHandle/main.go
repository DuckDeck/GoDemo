package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	// "path"
	// "path/filepath"
	// "strconv"
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
	r.POST("/upload", uploadHandler)
	r.Run(":9090")
}

func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	fmt.Println(header)
	if err != nil {
		//ignore
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {

	}
}
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
