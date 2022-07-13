package dto

type GetOrderListResponse struct {
	PageNum   int             `json:"page"`
	PageSize  int             `json:"page_size"`
	OrderList []OrderResponse `json:"order_list"`
}

type OrderResponse struct {
	ID              int               `json:"id"`
	TransactionCode string            `json:"transaction_code"`
	TotalPrice      float64           `json:"total_price"`
	OrderStatus     string            `json:"order_status"`
	Address         string            `json:"address"`
	Phone           string            `json:"phone"`
	CustomerID      int               `json:"customer_id"`
	ProductList     []ProductResponse `json:"product_list"`
}
