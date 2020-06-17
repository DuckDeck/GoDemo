package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//因为数据库也能查好几个所以不需要 这东西了， 一下子脑子不好用了
// var wg sync.WaitGroup

var db *sql.DB //一个连接池
// var sqlRes = make(chan fiveCode, 10)

func main() {
	initDB()
	initWebService()
}
func initDB() (err error) {
	con := "root:Stanhu520.@tcp(139.186.66.34:3306)/Five" //这个要放外面
	db, err = sql.Open("mysql", con)                      //已经有全局DB了，err也已经有了，不需要声明，这可能有很多坑
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

func checkExist(letters []string) (fives []FiveCode, err error) {
	var sql = fmt.Sprintf(`select id,word,pin_yin,five_code,img_code,img_keyboard from five_code where word in ('%s')`, strings.Join(letters, ","))
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("query fail errL", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var five FiveCode
		err = rows.Scan(&five.ID, &five.Word, &five.PinYin, &five.FiveCode, &five.ImgCode, &five.ImgKeyboard)
		if err != nil {
			fmt.Println("scan fail", err)
			return
		}
		fives = append(fives, five)
	}
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

func sendGet(key string) (fives []FiveCode, err error) {
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
		var item = FiveCode{Word: a, PinYin: b, FiveCode: c, ImgCode: d}
		fives = append(fives, item)
		saveToMysql(item)
	})
	return
}

type FiveCode struct {
	ID          int
	Word        string
	PinYin      string
	FiveCode    string
	ImgCode     string
	ImgKeyboard string
}

func saveToMysql(item FiveCode) {
	sql := `insert into five_code(word,pin_yin,five_code,img_code,img_keyboard) values(?,?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Printf("sql: %s prepare fail :%a\n", sql, err)
		return
	}
	defer stmt.Close()
	//后续只要用到stmp
	res, err := stmt.Exec(item.Word, item.PinYin, item.FiveCode, item.ImgCode, item.ImgKeyboard)
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

func initWebService() {
	http.HandleFunc("/getfive", checkData)
	http.ListenAndServe("0.0.0.0:9000", nil)
}

func checkData(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	key := param.Get("key")
	reg := regexp.MustCompile("[\u4e00-\u9fa5]")
	res := reg.FindAllStringSubmatch(key, 20)
	if len(res) <= 0 {
		w.Write([]byte(createResult(0, "你没有输入汉字")))
		return
	}

	var arrLetters []string
	for _, p := range res {
		arrLetters = append(arrLetters, p[0])
	}

	var result string
	fives, err := checkExist(arrLetters)
	if err != nil {
		result = createResult(-100, err.Error())
		w.Write([]byte(result))
		return
	}
	fmt.Println(fives)
	jsStr, _ := json.Marshal(fives)
	fmt.Printf("json:%s\n", string(jsStr))
	w.Write([]byte(createResult(0, string(jsStr))))
}

func createResult(code int, str string) (res string) {
	var result = Result{Code: code, Data: str}
	data, _ := json.Marshal(result)
	res = string(data)
	return
}

type Result struct {
	Code int
	Data string
}

func removeItem(arr []string, item string) []string {
	var index = 0
	for i, s := range arr {
		if s == item {
			index = i
		}
	}
	//会改变自己，所以要用copy
	tmp := make([]string, len(arr), len(arr))
	copy(tmp, arr)
	return append(tmp[:index], tmp[(index+1):]...)
}
