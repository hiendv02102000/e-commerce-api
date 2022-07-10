package entity

type orderStatus string

const (
	NOT_ORDERED orderStatus = "not ordered"
	ORDERED     orderStatus = "ordered"
	SHIPPING    orderStatus = "shipping"
	CANCLED     orderStatus = "canceled"
	COMPLETED   orderStatus = "completed"
	REJECTED    orderStatus = "rejected"
)

// Order struct
type Order struct {
	ID                  int               `gorm:"column:id;primary_key;auto_increment;not null"`
	TransactionCode     string            `gorm:"column:transaction_code;not null;unique;type:varchar(255)"`
	TotalPrice          float64           `gorm:"column:total_price;not null"`
	OrderStatus         orderStatus       `gorm:"column:order_status;not null"`
	Address             string            `gorm:"column:address;not null;type:varchar(255)"`
	CustomerID          int               `gorm:"column:customer_id;not null"`
	Customer            Users             `gorm:"foreignKey:customer_id;references:id"`
	SelectedProductList []SelectedProduct `gorm:"foreignKey:order_id;references:id"`
	BaseModelWithDeleteAt
}

func (c *Order) TableName() string {
	return "orders"
}
