package controller

import (
	"net/http"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/interactor"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUser(c *gin.Context) error
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
	//fmt.Println("[[[[]]]]", body.Email, "[[[[]]]]")

	u, err := uc.userInteractor.Get(u)
	if err != nil {

		return err
	}
	c.JSON(http.StatusCreated, u)
	return nil
}
