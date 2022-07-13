package dto

type CancelOrderRequest struct {
	TransactionCode string `json:"transaction_code" form:"transaction_code" binding:"required"`
}
