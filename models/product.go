package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code        string
	Name        string
	Description string
	Category    string
	Tags        pq.StringArray `gorm:"type:text[]"`
	Status      int
	Amount      int
	Price       float64
	CreatedAt   time.Time
	CreatedBy   uuid.UUID `gorm:"type:uuid"`
}
