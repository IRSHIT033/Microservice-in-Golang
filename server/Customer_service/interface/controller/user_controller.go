package controller

import (
	"net/http"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/interactor"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUser(c *gin.Context) error
	CreateUser(c *gin.Context) error
	AddToCart(c *gin.Context) error
	GetCart(c *gin.Context) error
	RemoveFromCart(c *gin.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUser(c *gin.Context) error {
	var u *model.User

	//bind body with the context
	if c.Bind(&u) != nil {
		helper.ShowError(c, "failed to read the body")
		return nil
	}

	u, err := uc.userInteractor.Get(u)
	if err != nil {

		return err
	}
	c.JSON(http.StatusOK, u)
	return nil
}

func (uc *userController) CreateUser(c *gin.Context) error {
	var user *model.User
	//bind body with the context
	if c.Bind(&user) != nil {
		helper.ShowError(c, "failed to read the body")
		return nil
	}
	//Generate Hash from the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		helper.ShowError(c, "failed to generate hash")
		return nil
	}

	//put hash into password field
	user.Password = string(hash)

	user, err = uc.userInteractor.Create(user)

	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, user)
	return nil
}

func (uc *userController) AddToCart(c *gin.Context) error {
	var product *model.Product

	//bind body with the context
	if c.Bind(&product) != nil {
		helper.ShowError(c, "failed to read the body")
		return nil
	}

	msg, err := uc.userInteractor.AddProductToCustomersCart(product)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, msg)

	return nil
}

func (uc *userController) GetCart(c *gin.Context) error {
	//get id from body param
	customerId := c.Param("id")
	customerID, _ := strconv.ParseUint(customerId, 10, 32)
	var product []*model.Product
	product, err := uc.userInteractor.GetProductinCustomersCart(uint(customerID))

	if err != nil {
		return err
	}
	c.JSON(http.StatusFound, product)

	return nil
}

func (uc *userController) RemoveFromCart(c *gin.Context) error {
	var body struct {
		CustomerID uint
		ProductID  uint
	}

	//bind body with the context
	if c.Bind(&body) != nil {
		helper.ShowError(c, "failed to read the body")
		return nil
	}

	msg, err := uc.userInteractor.RemoveProductFromCustomersCart(body.CustomerID, body.ProductID)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, msg)

	return nil

}
