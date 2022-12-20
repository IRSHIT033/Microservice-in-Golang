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
	Addresses []Address
	Cart      []Product
	Wishlist  []WishlistOfUser
	Orders    []Order
}
type Address struct {
	Street     string
	PostalCode string
	Country    string
}

type Product struct {
	ID    string
	Name  string
	Price string
	unit  int
}

type WishlistOfUser struct {
	ID          string
	Name        string
	Description string
	Available   bool
	price       int
}

type Order struct {
	ID     string
	Amount string
	Date   time.Time
}
