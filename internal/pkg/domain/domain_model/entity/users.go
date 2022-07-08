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
	Email          string     `gorm:"column:email;not null;unique"`
	Password       string     `gorm:"column:password;not null"`
	FirstName      string     `gorm:"column:first_name;"`
	LastName       string     `gorm:"column:last_name"`
	Role           userRole   `gorm:"column:role"`
	AvatarUrl      *string    `gorm:"column:avatar_url"`
	Token          *string    `gorm:"column:token"`
	TokenExpiredAt *time.Time `gorm:"column:token_expired_at"`
	BaseModel
}

func (u *Users) TableName() string {
	return "users"
}
