package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json::"status"`
}

var db *gorm.DB

func initDB() (err error) {

	db, err = gorm.Open("mysql", "root:Stanhu520.@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	return db.DB().Ping()
}
func main() {

	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "Project/Bubble/static")
	r.LoadHTMLGlob("Project/Bubble/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			if err = db.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})

		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			if err = db.Find(&todos).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todos)
			}

		})

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
			}
			var todo Todo
			if err = db.Where("id = ?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			c.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
			}
			if err = db.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})

			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "success"})
			}
		})
	}

	r.Run()
}
