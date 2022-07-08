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
	ID              int         `gorm:"column:id;primary_key;auto_increment;not null"`
	TransactionCode string      `gorm:"column:transaction_code;not null;unique"`
	TotalPrice      float64     `gorm:"column:total_price;not null"`
	OrderStatus     orderStatus `gorm:"column:order_status;not null"`
	CustomerID      int         `gorm:"column:customer_id;not null"`
	Customer        Users       `gorm:"foreignKey:CustomerID"`
	BaseModelWithDeleteAt
}

func (c *Order) TableName() string {
	return "orders"
}
