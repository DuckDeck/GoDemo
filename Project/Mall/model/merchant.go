package model

import (
	"github.com/jinzhu/gorm"
)

type Merchant struct {
	gorm.Model
	Logo   string
	Poster string
	Name   string
	SalNum uint
}
