package entity

// Brand struct
type Brand struct {
	ID    int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Title string `gorm:"column:title;not null;type:varchar(255)"`
	BaseModelWithDeleteAt
}

func (br *Brand) TableName() string {
	return "brands"
}
