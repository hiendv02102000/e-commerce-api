package dto

type GetOrderListRequest struct {
	PageNum         int `form:"page" binding:"omitempty"`
	CustomerID      int
	TransactionCode string `form:"transaction_code" binding:"omitempty"`
	Phone           string `form:"phone" binding:"omitempty"`
	Address         string `form:"address" binding:"omitempty"`
	OrderStatus     string `form:"order_status" binding:"omitempty,oneof=shipping rejected completed ordered"`
}
