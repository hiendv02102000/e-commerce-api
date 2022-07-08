package entity

// SelectedProduct struct
type SelectedProduct struct {
	ID       int     `gorm:"column:id;primary_key;auto_increment;not null"`
	Price    float64 `gorm:"column:price;not null"`
	Quantity int     `gorm:"column:quantity;not null"`
	BaseModel
}

func (p *SelectedProduct) TableName() string {
	return "selected_products"
}
