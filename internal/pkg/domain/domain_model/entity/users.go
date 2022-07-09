package entity

import (
	"time"
)

type userRole string

const (
	AdminRole userRole = "admin"
	Customer  userRole = "customer"
)

// Users struct
type Users struct {
	ID             int        `gorm:"column:id;primary_key;auto_increment;not null"`
	Email          string     `gorm:"column:email;not null;unique;type:varchar(255)"`
	Password       string     `gorm:"column:password;not null;type:varchar(255)"`
	FirstName      string     `gorm:"column:first_name;type:varchar(255)"`
	LastName       string     `gorm:"column:last_name;type:varchar(255)"`
	Role           userRole   `gorm:"column:role;type:varchar(255)"`
	AvatarUrl      *string    `gorm:"column:avatar_url;type:varchar(255)"`
	Token          *string    `gorm:"column:token;type:varchar(255)"`
	TokenExpiredAt *time.Time `gorm:"column:token_expired_at;type:varchar(255)"`
	BaseModel
}

func (u *Users) TableName() string {
	return "users"
}
