package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionRequestBody struct {
	PaidAmount                   float64                        `json:"paid_amount"`
	TransactionDetailRequestBody []TransactionDetailRequestBody `json:"transaction_detail"`
}

type TransactionResponse struct {
	ID                 uuid.UUID                   `json:"id"`
	TotalPrice         float64                     `json:"total_price"`
	IsSuccessPaid      bool                        `json:"is_success_paid"`
	CreatedAt          time.Time                   `json:"created_at"`
	TransactionDetails []TransactionDetailResponse `json:"transaction_details"`
}
