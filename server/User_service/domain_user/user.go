package domain_user

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Name      string
	Password  string
	Addresses []Address `gorm:"foreignKey:Belongsto"`
	Cart      []Product `gorm:"many2many:Products_In_Cart;"`
}

type Address struct {
	gorm.Model
	Belongsto  uint
	Street     string
	PostalCode string
	Country    string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, int, error)
}
