package model

import "github.com/jinzhu/gorm"

type Banner struct {
	gorm.Model

	Title  string `json:"title"`
	Img    string `json:"img"`
	Height uint   `json:"height"`
	Wight  uint   `json:"wight"`
	Link   string `json:"link"`
}
