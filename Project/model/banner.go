package model

import (
	"GoDemo/Project/Mall/global"

	"github.com/jinzhu/gorm"
)

type Banner struct {
	gorm.Model

	Title  string `json:"title"`
	Img    string `json:"img"`
	Height int    `json:"height"`
	Wight  int    `json:"wight"`
	Link   string `json:"link"`
}

func GetTopBanner(top int) (banners []Banner, err error) {

	global.G_DB.Find(&banners)
	return
}
