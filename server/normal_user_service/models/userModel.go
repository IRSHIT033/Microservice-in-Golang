package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Name      string
	Password  string
	Addresses []Address        `gorm:"foreignKey:Belongsto"`
	Cart      []Product        `gorm:"foreignKey:ProductID"`
	Wishlist  []WishlistOfUser `gorm:"foreignKey:WishlistProductID"`
	Orders    []Order          `gorm:"foreignKey:OrderID"`
}
type Address struct {
	Belongsto  uint
	Street     string
	PostalCode string
	Country    string
}

type Product struct {
	gorm.Model
	ProductID uint
	Name      string
	Price     string
	unit      int
}

type WishlistOfUser struct {
	gorm.Model
	WishlistProductID uint
	Name              string
	Description       string
	Available         bool
	Price             int
}

type Order struct {
	gorm.Model
	OrderID uint
	Amount  string
	Date    time.Time
}
