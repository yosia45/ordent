package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	Basemodel
	TransactionID uuid.UUID `json:"transaction_id" gorm:"not null;size:191"`
	ItemID        uuid.UUID `json:"item_id" gorm:"not null;size:191"`
	Quantity      int       `json:"quantity" gorm:"not null"`
	PricePerUnit  float64   `json:"price_per_unit" gorm:"not null"`
	TotalPrice    float64   `json:"total_price" gorm:"not null"`
}

func (td *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	td.ID = uuid.New()
	td.CreatedAt = time.Now()

	return
}
