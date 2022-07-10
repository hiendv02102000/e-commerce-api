package dto

type CreateCategoryRequest struct {
	ID    int    `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}
