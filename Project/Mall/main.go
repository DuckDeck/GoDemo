package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"` //无视这个
}

func a(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello you"})
}
func startServer() {
	r := gin.Default()
	r.GET("/hello", a)
	r.POST("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.Run(":8889")
}
func initDB() {
	db, err := gorm.Open("mysql", "root:Stanhu520.@tcp(127.0.0.1:3306)/mall")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) //默认新建表都会用复数形式，使用这个来让表使用单数形式
	//db.AutoMigrate(&User{})

	//db.Table("test_table").CreateTable(&testTable{}) //一般只用一次
	t := testTable{Age: 11, Name: "avc"}
	fmt.Println(db.NewRecord(&t)) //对于一个对象来，如果本身被用于db.create了，
	//那么就相当于被使用了，所以就已经存在了，不过如果重新再跑一次，前面的内存清空了，这就是一个新的，这样的话可以再存进数据库一次
	db.Create(&t)
	fmt.Println(db.NewRecord(&t))
}
func main() {
	//startServer()
	initDB()
}

type testTable struct { //要用gorm建表的话一定要加上gorm.Model而且字段名要大写
	gorm.Model
	Age  int
	Name string
}
