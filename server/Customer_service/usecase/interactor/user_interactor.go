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
	Create(*model.User) (*model.User, error)
	AddProductToCustomersCart(*model.Product) (string, error)
	GetProductinCustomersCart(uint) ([]*model.Product, error)
	RemoveProductFromCustomersCart(uint, uint) (string, error)
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

func (us *userInteractor) Create(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Save(u)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUser(u), nil
}

func (us *userInteractor) AddProductToCustomersCart(product *model.Product) (string, error) {
	msg, err := us.UserRepository.AddProduct(product)
	if err != nil {
		return "Error in the uscase Interactor", err
	}
	return us.UserPresenter.ResponseWithMessage(msg), nil
}

func (us *userInteractor) GetProductinCustomersCart(userId uint) ([]*model.Product, error) {
	var products []*model.Product
	products, err := us.UserRepository.GetCart(userId)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseProducts(products), nil
}

func (us *userInteractor) RemoveProductFromCustomersCart(userID uint, productId uint) (string, error) {
	msg, err := us.UserRepository.RemoveProduct(userID, productId)
	if err != nil {
		return "Error in the uscase Interactor", err
	}
	return us.UserPresenter.ResponseWithMessage(msg), nil
}
