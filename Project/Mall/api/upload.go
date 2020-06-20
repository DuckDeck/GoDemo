package api

import (
	"GoDemo/Project/Mall/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	header, err := c.FormFile("upload-key")
	if err != nil {
		fmt.Println(err)
		var res = model.Result{Code: -11, Count: 0, Msg: "request not contail file", Data: nil}
		c.JSON(http.StatusOK, res)
		return
	}

	dst := strconv.FormatInt(time.Now().Unix(), 10) + header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, "static/img/"+dst); err != nil {
		fmt.Println(err)
		var res = model.Result{Code: -10, Count: 0, Msg: "save file fail", Data: nil}
		c.JSON(http.StatusOK, res)

	} else {
		var data = make(map[string]string, 1)
		data["data"] = "https://lovelive.ink:19996/file/" + dst
		var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: data}
		c.JSON(http.StatusOK, res)
	}
}
