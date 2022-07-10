package dto

type DeleteCategoryRequest struct {
	CategoryID int `json:"category_id" form:"category_id"`
}
