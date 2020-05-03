package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Images     []string
	UnitTitle  string
	UnitSku    []Sku
	SaleDesc   string
	SalNum     uint
	CollectNum uint
	ShopID     uuid.UUID
}

type GoodsCat struct {
	Name   string
	Gender string
}

type Sku struct {
	gorm.Model
	UnitID int
	Size   string
	Color  string
	Look   string
}

type UnitSize struct {
	gorm.Model
	cat  GoodsCat
	XXXS float32
	XXS  float32
	XS   float32
	S    float32
	M    float32
	L    float32
	XL   float32
	XXL  float32
	XXXL float32
}
