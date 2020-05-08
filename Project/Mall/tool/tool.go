package tool

import (
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"os"

	simplejson "github.com/bitly/go-simplejson"
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

func FillBanner() {
	res, err := http.Get("http://123.207.32.32:8000/api/m3/home/data?type=sell&page=4")
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	js, err := simplejson.NewJson(body)
	banners, err := js.Get("data").Get("banner").Get("list").Array()
	if err != nil {
		return
	}
	var db = global.G_DB
	for _, v := range banners {
		var item = model.Banner{}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["title"].(string); ok {
				item.Title = t
			}
		}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["image"].(string); ok {
				item.Img = t
			}
		}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["link"].(string); ok {
				item.Link = t
			}
		}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["height"].(json.Number); ok {
				var z, _ = t.Int64()

				item.Height = *(*int)(unsafe.Pointer(&z))
			}
		}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["width"].(json.Number); ok {
				var z, _ = t.Int64()
				item.Wight = *(*int)(unsafe.Pointer(&z))
			}
		}
		db.Create(&item)
	}
}

func FillGoods() {
	//http://123.207.32.32:8000/api/m3/home/data?type=pop&page=1
	res, err := http.Get("http://123.207.32.32:8000/api/m3/home/data?type=pop&page=5")
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	js, err := simplejson.NewJson(body)
	banners, err := js.Get("data").Get("list").Array()
	if err != nil {
		return
	}
	var db = global.G_DB
	for k, v := range banners {
		var item = model.Goods{}
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["sale"].(json.Number); ok {
				var z, _ = t.Int64()
				item.SalNum = *(*int)(unsafe.Pointer(&z))
			}
			if t, ok := each_map["itemType"].(json.Number); ok {
				var z, _ = t.Int64()
				item.GoodsCatID = *(*int)(unsafe.Pointer(&z))
			}
			if t, ok := each_map["title"].(string); ok {
				item.UnitTitle = t
			}
			if t, ok := each_map["link"].(string); ok {
				item.Link = t
			}
			if t, ok := each_map["orgPrice"].(string); ok {
				var tmp = strings.Replace(t, "￥", "", 1)
				var s, _ = strconv.ParseFloat(tmp, 2)
				item.OrgPrice = s
			}
			if t, ok := each_map["sale"].(json.Number); ok {
				item.OrgPrice, _ = t.Float64()
			}
			if t, ok := each_map["shopId"].(json.Number); ok {
				var z, _ = t.Int64()
				item.MerchantID = *(*int)(unsafe.Pointer(&z))
			}
			if t, ok := each_map["cfav"].(json.Number); ok {
				var z, _ = t.Int64()
				item.UnitID = *(*int)(unsafe.Pointer(&z))
			}
			if t, ok := each_map["show"].(map[string]interface{}); ok {
				item.MainImage = t["img"].(string)
			}

		}

		rand.Seed(time.Now().Unix() + int64(k))
		item.CollectNum = rand.Intn(10000)
		item.CheckNum = rand.Intn(10000)
		item.Status = 1
		item.Keyword = item.UnitTitle
		item.CreateTime = time.Now()

		db.Create(&item)
	}
}

func FillCat() {
	res, err := http.Get("http://123.207.32.32:8000/api/m3/category")
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	js, err := simplejson.NewJson(body)
	banners, err := js.Get("data").Get("list").Array()
	if err != nil {
		return
	}
	var db = global.G_DB
	for k, v := range banners {
		var item = model.GoodsCat{PID: 0, Level: 1}
		item.Sort = k
		if each_map, ok := v.(map[string]interface{}); ok {
			if t, ok := each_map["title"].(string); ok {
				item.Name = t
			}
		}
		item.FullPath = item.Name
		db.Create(&item)
	}
}
