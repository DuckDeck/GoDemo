package main

import (
	"database/sql"
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
	db.AutoMigrate(&User{})
}
func main() {
	//startServer()
	initDB()
}
