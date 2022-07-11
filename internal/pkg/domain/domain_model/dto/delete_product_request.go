package dto

type DeleteProductRequest struct {
	ProductID int `json:"product_id" form:"product_id"`
}
