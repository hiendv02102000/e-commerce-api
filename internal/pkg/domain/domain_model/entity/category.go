package entity

// Category struct
type Category struct {
	ID    int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Title string `gorm:"column:title;not null"`
	BaseModelWithDeleteAt
}

func (c *Category) TableName() string {
	return "categories"
}
