package main

import (
	"github.com/gin-gonic/gin"
)

func a(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello you"})
}

func main() {
	r := gin.Default()
	r.GET("/hello", a)
	r.POST("/b",func(c * gin.Context){
		c.JSON(200.gin.H{
			"method":"POST"
		})
	})
	r.Run(":8889")
}
