package usecase

import (
	"context"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/domain_product"
)

type productUsecase struct {
	productRepository domain_product.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain_product.ProductRepository, timeout time.Duration) domain_product.ProductRepository {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) Create(c context.Context, product domain_product.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.Create(ctx, product)
}

func (pu *productUsecase) Fetch(c context.Context) ([]domain_product.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.Fetch(ctx)
}

func (pu *productUsecase) FetchbyId(c context.Context, productID uint) (domain_product.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.FetchbyId(ctx, productID)
}

func (pu *productUsecase) FetchbyCategory(c context.Context, categories []string) ([]domain_product.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.FetchbyCategory(ctx, categories)
}
