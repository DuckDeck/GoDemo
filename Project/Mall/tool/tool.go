package tool

import (
	"GoDemo/Project/Mall/model"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//对于get参数，要么是int，要么是string 再加个datetime
func CheckNum(c *gin.Context, paraName string) (res model.Result) {
	var para, yes = c.Get(paraName)
	if !yes {
		res.Code = -100
		res.MsgDebug = "miss para:" + paraName
		return
	}
	//不知道会设定成string还是int
	switch para.(type) {
	case string:
		fmt.Println("string", para.(string))
	case int:
		fmt.Println("string", para.(string))
	case float64:
		fmt.Println("string", para.(string))
	}
	res.Code = 0
	return
}

func CheckPage(c *gin.Context) (res model.Result) {
	var para1, yes = c.Get("pagesize")
	if !yes {
		res.Code = -100
		res.MsgDebug = "miss para: pagesize"
		return
	}
	//不知道会设定成string还是int
	switch para1.(type) {
	case string:
		fmt.Println("string", para1.(string))
	case int:
		fmt.Println("string", para1.(string))
	case float64:
		fmt.Println("string", para1.(string))
	}
	res.Code = 0
	return
}
