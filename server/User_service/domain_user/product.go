package domain_user

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	AddedBy         uint
	ProductID       uint
	ProductImageSrc string
	Name            string
	Description     string
	Price           float32
	Unit            int
	Available       bool
	Category        string
}

type Productrepository interface {
	Add(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID uint) ([]Product, error)
	FetchUserCart(c context.Context, userID uint) (User, error)
	Remove(c context.Context, userID uint, productID uint) error
}

type ProductUseCase interface {
	Add(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID uint) ([]Product, error)
	FetchUserCart(c context.Context, userID uint) (User, error)
	Remove(c context.Context, productID uint, userID uint) error
}
