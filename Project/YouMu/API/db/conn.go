package db

import (
	"database/sql"
	"fmt"
)

var (
	db  *sql.DB
	err error
)

func init() {
	con := "root:Stanhu520.@tcp(127.0.0.1:3306)/YouMu" //这个要放外面
	var err error
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
