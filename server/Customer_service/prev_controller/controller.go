package controller

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/controller/user"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/repository"
)

// Controller contains all the Controllers
type Controllers struct {
	userController *user.Controller
}

//InitControllers return a new Controller

func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		userController: user.InitController(repositories.UserRepo),
	}

}
