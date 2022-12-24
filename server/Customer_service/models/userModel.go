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
	Cart      []Product        `gorm:"foreignKey:AddedBy"`
	Wishlist  []WishlistOfUser `gorm:"foreignKey:WishlistBelongsto"`
	Orders    []Order          `gorm:"foreignKey:OrderID"`
}

type Address struct {
	gorm.Model
	Belongsto  uint
	Street     string
	PostalCode string
	Country    string
}

type Product struct {
	gorm.Model
	AddedBy         uint
	ProductID       uint
	ProductImageSrc string
	Name            string
	Price           int
	Unit            int
}

type WishlistOfUser struct {
	gorm.Model
	ProductNumber     uint
	ProductImageSrc   string
	WishlistBelongsto uint
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
