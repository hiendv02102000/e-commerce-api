package dto

type CreateBrandRequest struct {
	Title string `json:"title" form:"title" binding:"required"`
}
