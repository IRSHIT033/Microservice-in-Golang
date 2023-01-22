package usecase

import (
	"context"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/domain_order"
)

type orderUsecase struct {
	orderRepository domain_order.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUsecase(orderRepository domain_order.OrderRepository, timeout time.Duration) domain_order.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) Create(c context.Context, CustomerId uint, transactionId uint) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Create(ctx, CustomerId, transactionId)
}
func (ou *orderUsecase) Fetch(c context.Context, CustomerId uint) ([]domain_order.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Fetch(ctx, CustomerId)
}
