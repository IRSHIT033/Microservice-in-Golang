package presenter

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUser(us *model.User) *model.User {

	return us
}

func (up *userPresenter) ResponseWithMessage(msg string) string {

	return msg
}

func (up *userPresenter) ResponseProducts(products []*model.Product) []*model.Product {
	return products
}
