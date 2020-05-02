package model

import (
	"github.com/gin-gonic/gin"
)

type ShopEvaluate struct {
	gin.Model
	EvaluateType int
	Poster       string
	Name         string
	SalNum       uint
}

type UnitEvaluate struct {
	gin.Model
	UnitId   int
	SenderID int
	Time     time
	Comment  string
	value    int
	UnitSku  Sku
}
