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
	Price           int
	Unit            int
}

type Productrepository interface {
	Add(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID uint) ([]Product, error)
	Remove(c context.Context, userID uint, productID uint) error
}

type ProductUseCase interface {
	Add(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID uint) ([]Product, error)
	Remove(c context.Context, product *Product) error
}
