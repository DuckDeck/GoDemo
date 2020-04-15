package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //一个连接池

func initDB() (err error) {
	con := "root:Stanhu520.@tcp(127.0.0.1:3306)/Five"
	db, err = sql.Open("mysql", con) //已经有全局DB了，err也已经有了，不需要声明，这可能有很多坑
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
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init DB fail err:", err)
		return
	}
	query()
	defer db.Close()
}

type fiveCode struct {
	id          int
	word        string
	pinYin      string
	fiveCode    string
	imgCode     sql.NullString
	imgKeyboard sql.NullString
}

func queryRow(id int) {
	sql := `select * from five_code where id = ?`
	row := db.QueryRow(sql, id)
	var five fiveCode
	err := row.Scan(&five.id, &five.word, &five.pinYin, &five.fiveCode, &five.imgCode, &five.imgKeyboard)
	if err != nil {
		fmt.Printf("sql: %s query fail :%a\n", sql, err)
		return
	}
	fmt.Println(five)
}

func query() {
	sql := `select * from five_code`
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("sql: %s query fail :%a\n", sql, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var five fiveCode
		err = rows.Scan(&five.id, &five.word, &five.pinYin, &five.fiveCode, &five.imgCode, &five.imgKeyboard)
		if err != nil {
			fmt.Println("scan fail", err)
			return
		}
		fmt.Println(five)
	}
}

func insert() {
	sql := `insert into five_code(word,pin_yin,five_code) values("为","wei","o")`
	ret, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("sql: %s query fail :%a\n", sql, err)
		return
	}
	var res int64
	res, err = ret.LastInsertId()
	if err != nil {
		fmt.Printf("get insert result fail :%a\n", err)
		return
	}
	fmt.Println("inser success thr id is ", res)
}
