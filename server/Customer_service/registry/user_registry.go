package registry

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/interface/controller"
	ip "github.com/IRSHIT033/E-comm-GO-/server/Customer_service/interface/presenter"
	ir "github.com/IRSHIT033/E-comm-GO-/server/Customer_service/interface/repository"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/interactor"
	up "github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/presenter"
	ur "github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())

}

func (r *registry) NewUserRepository() ur.UserRepository {

	return ir.NewUserRepository(r.db)

}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
