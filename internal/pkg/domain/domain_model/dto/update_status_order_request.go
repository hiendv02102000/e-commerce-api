package dto

type UpdateStatucOrderRequest struct {
	TransactionCode string `form:"transaction_code" binding:"omitempty"`
	OrderStatus     string `form:"order_status" binding:"oneof=shipping rejected completed"`
}
