package domain_order

import (
	"context"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId    uint
	Amount        int
	Status        string
	TransactionId uint
	Products      []Product `gorm:"foreignKey:AddedBy"`
}

type Cart struct {
	gorm.Model
	CustomerId uint
	Items      ItemList
}

type ItemList struct {
	Products []Product `gorm:"foreignKey:AddedBy"`
	Unit     int
}

type Product struct {
	gorm.Model
	AddedBy         uint
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
