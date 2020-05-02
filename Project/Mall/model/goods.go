package model

import "github.com/gin-gonic/gin"

type Goods struct {
	gin.Model
	Images     []string
	UnitTitle  string
	UnitSku    []Sku
	SaleDesc   string
	SalNum     uint
	CollectNum uint
}

type Sku struct {
	gin.Model
	UnitID int
	Size   string
	Color  string
	Look   string
}
