package interactor

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/presenter"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(*model.User) (*model.User, error)
}

func NewUserInteractor(repo repository.UserRepository, pres presenter.UserPresenter) UserInteractor {
	return &userInteractor{repo, pres}
}

func (us *userInteractor) Get(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Find(u)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUser(u), nil
}
