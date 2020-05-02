package model

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	Images     []string
	UnitTitle  string
	UnitSku    []Sku
	SaleDesc   string
	SalNum     uint
	CollectNum uint
}

type Sku struct {
	gorm.Model
	UnitID int
	Size   string
	Color  string
	Look   string
}
