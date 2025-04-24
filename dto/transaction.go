package dto

import "time"

type TransactionRequestBody struct {
	PaidAmount                   float64                        `json:"paid_amount"`
	TransactionDetailRequestBody []TransactionDetailRequestBody `json:"transaction_detail"`
}

type TransactionResponse struct {
	ID                 string                      `json:"id"`
	TotalPrice         float64                     `json:"total_price"`
	IsSuccessPaid      bool                        `json:"is_success_paid"`
	CreatedAt          time.Time                   `json:"created_at"`
	TransactionDetails []TransactionDetailResponse `json:"transaction_details"`
}
