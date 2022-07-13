package entity

type OrderStatus string

const (
	CART      OrderStatus = "cart"
	ORDERED   OrderStatus = "ordered"
	SHIPPING  OrderStatus = "shipping"
	CANCLED   OrderStatus = "canceled"
	COMPLETED OrderStatus = "completed"
	REJECTED  OrderStatus = "rejected"
)

// Order struct
type Order struct {
	ID                  int               `gorm:"column:id;primary_key;auto_increment;not null"`
	TransactionCode     string            `gorm:"column:transaction_code;not null;unique;type:varchar(255)"`
	TotalPrice          float64           `gorm:"column:total_price;not null"`
	Status              OrderStatus       `gorm:"column:order_status;not null"`
	Address             string            `gorm:"column:address;not null;type:varchar(255)"`
	Phone               string            `gorm:"column:phone;not null;type:varchar(13)"`
	CustomerID          int               `gorm:"column:customer_id;not null"`
	Customer            Users             `gorm:"foreignKey:customer_id;references:id"`
	SelectedProductList []SelectedProduct `gorm:"foreignKey:order_id;references:id"`
	BaseModelWithDeleteAt
}

func (c *Order) CalculatePrice() {
	c.TotalPrice = 0.0
	for _, product := range c.SelectedProductList {
		c.TotalPrice += product.Price * float64(product.Quantity)
	}
}
func (c *Order) TableName() string {
	return "orders"
}
