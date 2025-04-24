package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Basemodel
	TotalPrice         float64             `json:"total_price" gorm:"not null"`
	IsSuccessPaid      bool                `json:"is_success_paid" gorm:"default:false"`
	UserID             uuid.UUID           `json:"user_id" gorm:"not null;size:191"`
	TransactionDetails []TransactionDetail `json:"transaction_details" gorm:"foreignKey:TransactionID"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.CreatedAt = time.Now()

	return
}
