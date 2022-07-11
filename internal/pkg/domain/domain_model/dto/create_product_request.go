package dto

type CreateProductRequest struct {
	Name        string  `json:"name" form:"name" binding:"required"`
	Price       float64 `json:"price" form:"price" binding:"required"`
	Discount    float64 `json:"discount" form:"discount" binding:"omitempty"`
	Quantity    int     `json:"quantity" form:"quantity" binding:"omitempty"`
	Description string  `json:"description" form:"description" binding:"required"`
	BrandID     int     `json:"brand_id" form:"brand_id" binding:"required"`
	CategoryID  int     `json:"category_id" form:"category_id" binding:"required"`
}
