package entity

// SelectedProduct struct
type SelectedProduct struct {
	ID          int     `gorm:"column:id;primary_key;auto_increment;not null"`
	Price       float64 `gorm:"column:price;not null"`
	Quantity    int     `gorm:"column:quantity;not null"`
	OrderID     int     `gorm:"column:order_id;not null"`
	ProductID   int     `gorm:"column:product_id;not null"`
	ProductInfo Product `gorm:"foreignKey:product_id;references:id"`
	BaseModel
}

func (p *SelectedProduct) TableName() string {
	return "selected_products"
}
