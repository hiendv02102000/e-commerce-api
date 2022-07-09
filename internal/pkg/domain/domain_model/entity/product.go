package entity

// Product struct
type Product struct {
	ID              int      `gorm:"column:id;primary_key;auto_increment;not null"`
	Name            string   `gorm:"column:name;not null;type:varchar(255)"`
	Price           float64  `gorm:"column:price;not null"`
	Discount        float64  `gorm:"column:discount;not null"`
	Quantity        int      `gorm:"column:quantity;not null"`
	Description     string   `gorm:"column:description;not null"`
	ImgURL          string   `gorm:"column:img_url;not null;type:varchar(255)"`
	CategoryID      int      `gorm:"column:category_id;not null"`
	CategoryProduct Category `gorm:"foreignKey:category_id;references:id"`
	BrandID         int      `gorm:"column:brand_id;not null"`
	BrandProduct    Brand    `gorm:"foreignKey:brand_id;references:id"`

	BaseModelWithDeleteAt
}

func (p *Product) TableName() string {
	return "products"
}
