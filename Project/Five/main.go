package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var db *sql.DB //一个连接池
func main() {
	r, err := regexp.Compile("[\u4e00-\u9fa5]")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := r.FindAllStringSubmatch("我r你fffs", 5)
	var keys []string
	for _, p := range res {
		keys = append(keys, p[0]) //拼成一个数组
	}

	// var s = sql.NullString{String: "123", Valid: true}
	// fmt.Println(s)
	checkExist(keys)
	//initDB()
	//sendGet("他地")
}
func initDB() (err error) {
	con := "root:Stanhu520.@tcp(127.0.0.1:3306)/Five" //这个要放外面
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

func checkExist(letters []string) {

	for _, item := range letters {
		fmt.Println(item)
	}
}

func sendGet(key string) {
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
		var item = fiveCode{word: a, pinYin: b, fiveCode: c, imgCode: sql.NullString{String: d, Valid: true}}
		saveToMysql(item)
	})

}

type fiveCode struct {
	id          int
	word        string
	pinYin      string
	fiveCode    string
	imgCode     sql.NullString
	imgKeyboard sql.NullString
}

func saveToMysql(item fiveCode) {
	sql := `insert into five_code(word,pin_yin,five_code,img_code) values(?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Printf("sql: %s prepare fail :%a\n", sql, err)
		return
	}
	defer stmt.Close()
	//后续只要用到stmp
	res, _ := stmt.Exec(item.word, item.pinYin, item.fiveCode, item.imgCode)
	id, _ := res.LastInsertId()
	fmt.Println("inser success thr id is ", id)
}

func sendFiveCode() {

}
