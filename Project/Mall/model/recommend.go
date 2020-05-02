package model

import "github.com/jinzhu/gorm"

type Recommend struct {
	gorm.Model
	Title string `json:"title"`
	Img   string `json:"img"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
}
