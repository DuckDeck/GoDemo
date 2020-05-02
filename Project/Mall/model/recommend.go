package model

import "github.com/gin-gonic/gin"

type Recommend struct {
	gin.Model
	Title string `json:"title"`
	Img   string `json:"img"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
}
