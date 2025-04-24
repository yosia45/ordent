package dto

type TransactionDetailRequestBody struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type TransactionDetailResponse struct {
	Item         GetItemDetailTransactionResponse `json:"item"`
	Quantity     int                              `json:"quantity"`
	PricePerUnit float64                          `json:"price_per_unit"`
	TotalPrice   float64                          `json:"total_price"`
}
