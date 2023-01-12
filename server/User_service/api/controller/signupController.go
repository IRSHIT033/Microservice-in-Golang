package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain_user.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {

	var request domain_user.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	_, userExists, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusConflict, domain_user.ErrorResponse{Message: "Error occured in usecase"})
		return
	}

	if userExists > 0 {
		c.JSON(http.StatusNotFound, domain_user.ErrorResponse{Message: "user already exists "})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain_user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}
	secret_key := os.Getenv("SECRET_KEY")
	expiry_time := os.Getenv("EXPIRY_TIME")
	expiry, _ := strconv.Atoi(expiry_time)
	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, secret_key, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, secret_key, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain_user.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
