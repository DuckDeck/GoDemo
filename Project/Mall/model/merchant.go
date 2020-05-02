package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Merchant struct {
	gorm.Model
	UUID   uuid.UUID `json:"uuid"`
	Logo   string
	Poster string
	Name   string
	SalNum uint
}
