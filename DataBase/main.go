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
	transaction()
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

func update() {
	sql := `update five_code set img_code = "https://www.52wubi.com/wbbmcx/tp/%E4%B8%BA.gif" where word = '为'`
	ret, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("sql: %s query fail :%a\n", sql, err)
		return
	}
	var res int64
	res, err = ret.RowsAffected()
	if err != nil {
		fmt.Printf("get insert result fail :%a\n", err)
		return
	}
	fmt.Println("alter RowsAffected is ", res)
}

//delete就不需要写了
//mysql
func prepareInsert() {
	sql := `insert into five_code(word,pin_yin,five_code) values(?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Printf("sql: %s prepare fail :%a\n", sql, err)
		return
	}
	defer stmt.Close()
	//后续只要用到stmp
	res, _ := stmt.Exec("小", "xiao", "ih")
	id, _ := res.LastInsertId()
	fmt.Println("inser success thr id is ", id)
}

//事务
func transaction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin transcation fail err:", err)
		return
	}
	sql1 := `update five_code set img_keyboard = 'asdfadfs' where id = 1`
	sql2 := `update five_code set img_keyboard = 'zzzzzzzzzzzzzzz' where id = 2`
	_, err = tx.Exec(sql1)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql1出错了，回滚,err:", err)
		return
	}
	tx.Exec(sql2)
	_, err = tx.Exec(sql2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2出错了，回滚,err:", err)
		return
	}
	//上面两次事务执行成功，可以提交
	err = tx.Commit()
	if err != nil {
		fmt.Println("事务提交出错回滚 err:", err)
		return
	}
	fmt.Println("事务执行成功")

}
