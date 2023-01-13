package usecase

import (
	"context"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
)

type productUsecase struct {
	productRepository domain_user.Productrepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain_user.Productrepository, timeout time.Duration) domain_user.Productrepository {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) Add(c context.Context, product *domain_user.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.Add(ctx, product)
}

func (pu *productUsecase) Remove(c context.Context, productID uint, userID uint) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.Remove(ctx, productID, userID)
}

func (pu *productUsecase) FetchByUserID(c context.Context, userID uint) ([]domain_user.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.FetchByUserID(ctx, userID)
}
