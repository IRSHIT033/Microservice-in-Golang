package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//read from the http body
	var body struct {
		Email    string
		Password string
		Name     string
	}
	if c.Bind(&body) != nil {
		helper.ShowError(c, "failed to read the body")
		return
	}
	// Genrate Hash from the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		helper.ShowError(c, "failed to generate hash")
		return
	}

	// create user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		helper.ShowError(c, "failed to create user")
		return
	}
	//respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {

	//get the email and password req body
	var body struct {
		Email    string
		Password string
		Name     string
	}
	if c.Bind(&body) != nil {
		helper.ShowError(c, "failed to read the body")
		return
	}
	//check if email registered

	var user models.User
	initializers.DB.First(&user, "email= ?", body.Email)
	if user.ID == 0 {
		helper.ShowError(c, "Invalid Email : email is not registered ")
		return
	}
	//check if the password and mail is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		helper.ShowError(c, "Invalid email or password")
		return
	}
	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Log in and the get the complete encoded token as a string using secret
	secret_key := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		helper.ShowError(c, "Failed to create token")
		return
	}
	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
