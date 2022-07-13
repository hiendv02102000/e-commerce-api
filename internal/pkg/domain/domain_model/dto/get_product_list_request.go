package dto

type GetProductListRequest struct {
	PageNum     int    `form:"page" binding:"omitempty"`
	BrandID     int    `form:"brand_id" binding:"omitempty"`
	CategoryID  int    `form:"category_id" binding:"omitempty"`
	ProductName string `form:"product_name" binding:"omitempty"`
}
