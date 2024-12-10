package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type InvoiceItem struct {
	Code     string
	Name     string
	Quantity int
	Price    float64
	Tax      float64
	Amount   float64
}

type Invoice struct {
	Id            uuid.UUID
	Code          string
	SellerName    string
	SellerAddress string
	BuyerName     string
	BuyerAddress  string
	Items         datatypes.JSON
	Total         float64
	Note          string
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
}
