package user

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/models"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/repository/userrepo"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type repository interface {
	GetExistingCustomer(email string, password string) models.User
	CreateCustomer(user models.User) models.User
	GetCartofCustomer(CustomerID uint) models.Product
	AddToCart(CustomerID uint, product models.Product)
	RemoveFromCart(CustomerID uint, ProductID uint)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	service repository
}

func InitController(userRepo *userrepo.UserRepo) *Controller {
	return &Controller{
		service: userRepo,
	}
}
