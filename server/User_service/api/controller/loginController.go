package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain_user.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain_user.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain_user.ErrorResponse{Message: err.Error()})
		return
	}
	user, userExists, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, domain_user.ErrorResponse{Message: "Error occured in usecase "})
		return
	}

	if userExists == 0 {
		c.JSON(http.StatusNotFound, domain_user.ErrorResponse{Message: "user not found with this given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain_user.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	secret_key := os.Getenv("SECRET_KEY")
	expiry_time := os.Getenv("EXPIRY_TIME")
	expiry, _ := strconv.Atoi(expiry_time)

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, secret_key, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, secret_key, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain_user.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)

}
