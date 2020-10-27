package model

import "github.com/jinzhu/gorm"

type Keyword struct {
	gorm.Model
	Word  string `json:"word"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
	IsRed bool   `json:"isRed"`
}
