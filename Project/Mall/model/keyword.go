package model

import "github.com/gin-gonic/gin"

type Keyword struct {
	gin.Model
	Word  string `json:"word"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
	IsRed bool   `json:"isRed"`
}
