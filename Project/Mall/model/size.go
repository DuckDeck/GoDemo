package model

import "github.com/gin-gonic/gin"

type Banner struct {
	gin.Model
	Title  string `json:"title"`
	Img    string `json:"img"`
	Height uint   `json:"height"`
	Wight  uint   `json:"wight"`
	Link   string `json:"link"`
}
