package dto

type UpdateProductRequest struct {
	ID          int     `json:"id" form:"id" binding:"required"`
	Name        string  `json:"name" form:"name" binding:"omitempty"`
	Price       float64 `json:"price" form:"price" binding:"omitempty"`
	Discount    float64 `json:"discount" form:"discount" binding:"omitempty"`
	Quantity    int     `json:"quantity" form:"quantity" binding:"omitempty"`
	Description string  `json:"description" form:"description" binding:"omitempty"`
	BrandID     int     `json:"brand_id" form:"brand_id" binding:"omitempty"`
	CategoryID  int     `json:"category_id" form:"category_id" binding:"omitempty"`
}
