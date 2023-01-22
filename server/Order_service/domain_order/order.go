package domain_order

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID       uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CustomerId    uint
	Amount        int
	Status        string
	TransactionId uint
	Products      []Product `gorm:"many2many:Products_In_Order;"`
}

type Cart struct {
	CartID     uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CustomerId uint
	Products   []Product `gorm:"many2many:Products_In_Cart;"`
	Unit       int
}

type Product struct {
	gorm.Model
	AddedIn         uint
	ProductID       uint
	ProductImageSrc string
	Name            string
	Description     string
	Price           int
	Unit            int
	Available       bool
	Category        string
}

type OrderRepository interface {
	Create(context context.Context, CustomerId uint, transactionId uint) error
	Fetch(context context.Context, CustomerId uint) ([]Order, error)
}

type OrderUsecase interface {
	Create(context context.Context, CustomerId uint, transactionId uint) error
	Fetch(context context.Context, CustomerId uint) ([]Order, error)
}
