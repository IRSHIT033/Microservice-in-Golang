package domain_product

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

type CategoryRequest struct {
	Category []string `json:"categories" binding:"required"`
}

type ProductRepository interface {
	Create(c context.Context, product Product) error
	Fetch(c context.Context) ([]Product, error)
	FetchbyId(c context.Context, productID uint) (Product, error)
	FetchbyCategory(c context.Context, category []string) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product Product) error
	Fetch(c context.Context) ([]Product, error)
	FetchbyId(c context.Context, productID uint) (Product, error)
	FetchbyCategory(c context.Context, category []string) ([]Product, error)
}
