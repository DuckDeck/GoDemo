package model

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Merchant struct {
	gin.Model
	UUID   uuid.UUID `json:"uuid"`
	Logo   string
	Poster string
	Name   string
	SalNum uint
}
