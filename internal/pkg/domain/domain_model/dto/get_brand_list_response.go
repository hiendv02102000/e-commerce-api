package dto

type GetBrandListResponse struct {
	Total     int             `json:"total"`
	BrandList []BrandResponse `json:"brand_list"`
}

type BrandResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
