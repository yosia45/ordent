package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Basemodel
	Name               string              `json:"name" gorm:"not null"`
	Price              float64             `json:"price" gorm:"not null"`
	Stock              int                 `json:"stock" gorm:"not null"`
	TransactionDetails []TransactionDetail `json:"transaction_details" gorm:"foreignKey:ItemID"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	i.CreatedAt = time.Now()

	return
}
