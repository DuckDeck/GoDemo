package api

import (
	"GoDemo/Project/Mall/model"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var imgPath = "/Download/Imgs"
var bathPath = "http://www.tuyiaa.com/"

func SaveMito(c *gin.Context) {
	path := c.Param("path")
	index, _ := strconv.Atoi(c.Param("index"))

	go getImgMain(bathPath+path, index)

	c.JSON(http.StatusOK, model.Result{Code: 0, Count: 0, Msg: "success", CMsg: "成功", Data: nil})
}

func getImgMain(baseUrl string, startIndex int) {
	var url string
	if startIndex < 0 {
		url = baseUrl
	} else {
		url = baseUrl + strconv.Itoa(startIndex) + ".html"
	}

	fmt.Println(url)
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
	var content, _ = doc.Html()
	var index = strings.Index(content, "btwaf")
	if index > 0 {
		var condiction = content[index+6 : index+14]
		var newUrl = url + "?btwaf=" + condiction
		getImgMain(newUrl, -1)
		return
	}
	pages := doc.Find("div.pg").First().Children().Length() + startIndex - 2
	pages = startIndex + 1
	fmt.Println("一共有多少页", pages)
	var arrCatImgs []string
	s := doc.Find("div.bus_vtem")
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		href, exist := s.Children().First().Attr("href")
		if !strings.HasPrefix(href, "http") {
			href = bathPath + href
		}
		if exist {
			arrCatImgs = append(arrCatImgs, href)
		}

	})
	for i := startIndex + 1; i < pages; i++ {
		res, err := http.Get(baseUrl + strconv.Itoa(i) + ".html")
		if err != nil {
			fmt.Println("request fail", err)
			return
		}
		reader := simplifiedchinese.GB18030.NewDecoder().Reader(res.Body) //需要将bgk转utf8
		doc, err = goquery.NewDocumentFromReader(reader)
		if err != nil {
			fmt.Println("解析HTML失败", err)
			return
		}

		s = doc.Find("div.bus_vtem")
		s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
			href, exist := s.Children().First().Attr("href")
			if exist {
				arrCatImgs = append(arrCatImgs, href)
			}
		})
	}
	for _, img := range arrCatImgs {
		getImgCat(img)
	}
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
	var content, _ = doc.Html()
	var index = strings.Index(content, "btwaf")
	if index > 0 {
		var condiction = content[index+6 : index+14]
		var newUrl = url + "?btwaf=" + condiction
		getImgCat(newUrl)
		return
	}

	title := doc.Find("#thread_subject").Text()
	cat, _ := doc.Find("a.bus_fl").Children().First().Attr("alt")
	mainImage, _ := doc.Find("ignore_js_op").First().Children().First().Attr("data-original")
	info, err := doc.Find("td.t_f").First().Children().Html()
	if err != nil {
		info = "无"
	}
	s := doc.Find(".savephotop")
	var arrImgs []string
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		img, exist := s.Children().First().Attr("data-original")
		if exist {
			arrImgs = append(arrImgs, img)
		}
	})
	if len(arrImgs) <= 0 {
		return
	}
	var tu = TuYi{urlStr: url, title: title, cat: cat, imgs: arrImgs, mainImage: mainImage, personInfo: info}

	var catImgPath = path.Join(imgPath, cat, title)

	os.MkdirAll(catImgPath, os.ModePerm)
	getImg(tu.mainImage, catImgPath)
	for _, img := range arrImgs {
		getImg(img, catImgPath)
	}
}

func getImg(img string, imgPath string) {
	fileName := path.Base(img)
	c := &http.Client{
		Timeout: 0,
	}
	res, err := c.Get(img)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 128*1024)
	path := filepath.Join(imgPath, fileName)
	fmt.Println(path)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}

type TuYi struct {
	urlStr     string
	title      string
	mainImage  string
	personInfo string
	cat        string
	imgs       []string
}
