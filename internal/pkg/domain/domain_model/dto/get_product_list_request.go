package dto

type GetProductListRequest struct {
	PageNum     int    `json:"page" form:"page" binding:"omitempty"`
	BrandID     int    `json:"brand_id" form:"brand_id" binding:"omitempty"`
	CategoryID  int    `json:"category_id" form:"category_id" binding:"omitempty"`
	ProductName string `json:"product_name" form:"product_name" binding:"omitempty"`
}
