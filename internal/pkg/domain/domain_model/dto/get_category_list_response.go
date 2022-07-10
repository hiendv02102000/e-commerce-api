package dto

type GetCategoryListResponse struct {
	Total        int                `json:"total"`
	CategoryList []CategoryResponse `json:"Category_list"`
}

type CategoryResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
