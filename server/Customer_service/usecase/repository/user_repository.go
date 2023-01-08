package repository

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
)

type UserRepository interface {
	Find(*model.User) (*model.User, error)
	// GetExistingCustomer(email string, password string) model.User
	// CreateCustomer(user model.User) model.User
	// GetCartofCustomer(CustomerID uint) model.Product
	// AddToCart(CustomerID uint, product model.Product)
	// RemoveFromCart(CustomerID uint, ProductID uint)
}
