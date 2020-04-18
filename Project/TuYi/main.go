package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"

	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var imgPath = "/Users/stan/Desktop/Project/GoDemo/Project/TuYi"

func main() {

}

func getImgCat(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("request fail", err)
		return
	}
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(res.Body) //需要将bgk转utf8
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("解析HTML失败", err)
		return
	}

	title := doc.Find("#thread_subject").Text()
	cat, _ := doc.Find("a.bus_fl").Children().First().Attr("alt")

	s := doc.Find(".savephotop")
	var arrImgs []string
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		img, exist := s.Children().First().Attr("data-original")
		if exist {
			arrImgs = append(arrImgs, img)
		}
	})

	var _ = TuYi{urlStr: url, title: title, cat: cat, imgs: arrImgs}
	for _, img := range arrImgs {
		getImg(img)

	}
}

func getImg(img string) {

	fileName := path.Base(img)
	res, err := http.Get(img)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 64*1024)

	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}

type TuYi struct {
	urlStr string
	title  string
	cat    string
	imgs   []string
}
