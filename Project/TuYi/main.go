package main

import (
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("http://www.tuyi8.vip/thread-26096-1-1.html")
	if err != nil {
		fmt.Println("request fail", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("解析HTML失败", err)
		return
	}
	s := doc.Find(".savephotop")
	var arrImgs []string
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		img, exist := s.Children().First().Attr("data-original")
		if exist {
			arrImgs = append(arrImgs, img)
		}
	})
}

type TuYi struct {
	urlStr string
	cat    string
	imgs   []string
}
