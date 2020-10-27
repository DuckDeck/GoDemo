package model

import (
	"GoDemo/Project/Mall/global"

	"github.com/jinzhu/gorm"
)

type Recommend struct {
	gorm.Model
	Title string `json:"title"`
	Img   string `json:"img"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
}

func GetRecommend(top int) (recommends []Recommend, err error) {
	global.G_DB.Find(&recommends)
	return
}
