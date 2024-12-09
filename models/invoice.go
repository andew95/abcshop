package models

import (
	"time"

	"github.com/google/uuid"
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
	Items         *[]Invoice
	Total         float64
	Note          string
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
}
