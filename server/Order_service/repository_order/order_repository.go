package repository_order

import (
	"context"
	"errors"
	"fmt"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/domain_order"
	"gorm.io/gorm"
)

type orderRepository struct {
	database *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain_order.OrderRepository {
	return &orderRepository{
		database: db,
	}
}

func (or *orderRepository) Create(c context.Context, CustomerId uint, transactionId uint) error {
	//Find cart
	var cart domain_order.Cart
	err := or.database.Where("customer_id = ?", CustomerId).Find(&cart).Error

	if cart.CustomerId == 0 {
		return errors.New("customer not found")
	}

	if err != nil {
		return err
	}
	//associate products in cart
	var products []domain_order.Product
	err = or.database.Model(&cart).Association("Products").Find(&products)
	if err != nil {
		return err
	}

	if len(products) == 0 {
		return errors.New("cart is empty")
	}

	//calculate amount
	var amount float32
	for _, item := range products {
		amount += item.Price
	}
	//create a order
	err = or.database.Create(&domain_order.Order{
		CustomerId:    cart.CustomerId,
		Amount:        amount,
		Status:        "received",
		TransactionId: transactionId,
		Products:      products,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) Fetch(c context.Context, CustomerId uint) ([]domain_order.Order, error) {
	var all_orders []domain_order.Order
	fmt.Println(CustomerId)
	or.database.Where("customer_id = ?", CustomerId).Preload("Products").Find(&all_orders)
	return all_orders, nil
}
