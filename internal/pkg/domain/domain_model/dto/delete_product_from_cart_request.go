package dto

type DeleteProductFromCartRequest struct {
	ProductID  int `json:"product_id" form:"product_id"`
	CustomerID int
}
