package dto

type PutProductToCartRequest struct {
	ProductID       int    `json:"product_id" binding:"required"`
	Quantity        int    `json:"quantity" binding:"min=1"`
	TransactionCode string `form:"transaction_code" binding:"required"`
}
