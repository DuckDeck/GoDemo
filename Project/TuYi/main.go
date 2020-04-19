package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var db *sql.DB //一个连接池
var imgPath = "/Users/stan/Desktop/Project/GoDemo/Project/TuYi/Imgs"

func main() {

	initDB()
	getImgMain("http://www.tuyi8.vip/forum-17-", 29)

}

func getImgMain(baseUrl string, startIndex int) {
	res, err := http.Get(baseUrl + strconv.Itoa(startIndex) + ".html")
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
	pages := doc.Find("div.pg").First().Children().Length() + startIndex - 2
	fmt.Println("一共有多少页", pages)
	var arrCatImgs []string
	s := doc.Find("div.bus_vtem")
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		href, exist := s.Children().First().Attr("href")
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
	saveToMysql(tu)
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

	file, err := os.Create(filepath.Join(imgPath, fileName))
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}

func saveToMysql(tu TuYi) {
	sql := `insert into beautiful(title,person_info,main_img,url,cat,imgs) values(?,?,?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Printf("sql: %s prepare fail :%a\n", sql, err)
		return
	}
	defer stmt.Close()
	//后续只要用到stmp
	res, err := stmt.Exec(tu.title, tu.personInfo, tu.mainImage, tu.urlStr, tu.cat, strings.Join(tu.imgs, "-"))
	if err != nil {
		fmt.Printf("sql: %s exec fail :%s\n", sql, err.Error())
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("sql: last ID fail :%s\n", err.Error())
		return
	}
	fmt.Println("inser success thr id is ", id)
}

func initDB() (err error) {

	con := "root:Stanhu520.@tcp(127.0.0.1:3306)/TuYi" //这个要放外面
	db, err = sql.Open("mysql", con)                  //已经有全局DB了，err也已经有了，不需要声明，这可能有很多坑
	if err != nil {
		fmt.Printf("conn %s invalid,err:%v\n", con, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s fail err:%v\n", con, err)
		return
	}
	fmt.Println("connect database success")
	db.SetMaxOpenConns(10) //最大连接数
	return
}

type TuYi struct {
	urlStr     string
	title      string
	mainImage  string
	personInfo string
	cat        string
	imgs       []string
}
