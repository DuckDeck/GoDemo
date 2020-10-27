package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ShopEvaluate struct {
	gorm.Model
	EvaluateType int
	Poster       string
	Name         string
	SalNum       uint
}

type UnitEvaluate struct {
	gorm.Model
	UnitId   int
	SenderID int
	Time     time.Time
	Comment  string
	value    int
	UnitSku  Sku
}
