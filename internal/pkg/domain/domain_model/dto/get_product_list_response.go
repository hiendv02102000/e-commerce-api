package dto

type GetProductListResponse struct {
	PageNum     int               `json:"page"`
	PageSize    int               `json:"page_size"`
	ProductList []ProductResponse `json:"product_list"`
}
type ProductResponse struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Price           float64          `json:"price"`
	Discount        float64          `json:"discount"`
	Quantity        int              `json:"quantity"`
	Description     string           `json:"description"`
	ImgURL          string           `json:"img_url"`
	BrandProduct    BrandResponse    `json:"brand"`
	CategoryProduct CategoryResponse `json:"category"`
}
