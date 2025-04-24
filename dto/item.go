package dto

import "github.com/google/uuid"

type ItemRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type GetAllItemResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
	Stock int       `json:"stock"`
}

type GetItemDetailTransactionResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
