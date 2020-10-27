package api

import (
	"GoDemo/Project/Mall/model"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetFive(c *gin.Context) {
	key := c.Param("key")
	reg := regexp.MustCompile("[\u4e00-\u9fa5]")
	res := reg.FindAllStringSubmatch(key, 20)
	if len(res) <= 0 {
		c.JSON(http.StatusOK, model.Result{Code: 123, Count: 0, Msg: "the search key must be chinese", CMsg: "搜索的字必定的汉字", Data: nil})
		return
	}

	var arrLetters []string
	for _, p := range res {
		arrLetters = append(arrLetters, p[0])
	}

	fives, err := checkExist(arrLetters)
	if err != nil {

		c.JSON(http.StatusOK, model.Result{Code: -100, Count: 0, Msg: err.Error(), CMsg: err.Error(), Data: nil})
		return
	}
	fmt.Println(fives)
	jsStr, _ := json.Marshal(fives)
	fmt.Printf("json:%s\n", string(jsStr))

	c.JSON(http.StatusOK, model.Result{Code: 0, Count: 0, Msg: "success", CMsg: "成功", Data: fives})
}

func checkExist(letters []string) (fives []model.FiveCode, err error) {
	fives, err = model.GetCodes(letters)
	if len(fives) == len(letters) {
		return
	}
	var arrRemaind []string
	for _, a := range letters {
		for _, b := range fives {
			if b.Word == a {
				continue
			}
		}
		arrRemaind = append(arrRemaind, a)
	}
	fives, err = sendGet(strings.Join(arrRemaind, ""))
	return
}

func sendGet(key string) (fives []model.FiveCode, err error) {
	client := &http.Client{}
	encoder := simplifiedchinese.GB18030.NewEncoder()
	newKey, _ := encoder.String(key)
	req, err := http.NewRequest("POST", "https://www.52wubi.com/wbbmcx/search.php", strings.NewReader("hzname="+newKey))
	if err != nil {
		fmt.Println("request fail", err)
		return
	}
	req.Header.Set("Accept", "text/html, application/xhtml+xml, image/jxr, */*")
	req.Header.Set("Referer", "https://www.52wubi.com/wbbmcx/search.php")
	req.Header.Set("Accept-Language", "zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("User-Agent", "tMozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "10")
	req.Header.Set("Proxy-Connection", "Keep-Alive")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request fail", err)
		return
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body) //需要将bgk转utf8
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println("解析HTML失败", err)
		return
	}
	s := doc.Find(".tbhover")
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		k := s.Children().First()
		a := k.Text()
		b := k.Next().Text()
		c := k.Next().Next().Text()
		d := k.Next().Next().Next().Children().First().AttrOr("src", "")
		d = "https://www.52wubi.com/wbbmcx/" + d
		fmt.Println(a, b, c, d)
		var item = model.FiveCode{Word: a, PinYin: b, FiveCode: c, ImgCode: d}
		fives = append(fives, item)
		model.SaveFive(item)
	})
	return
}
