package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/models"
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

func GetWishlist(c *gin.Context) {
	DB := initializers.DB
	//get id from param
	customerID := c.Param("id")
	//Check if the customer exists
	var user models.User
	profile := DB.First(&user, customerID)
	if profile.Error != nil {
		helper.ShowError(c, "Customer not found")
		return
	}
	//get the wishlist data of the customer
	DB.Model(&user).Association("Wishlist").Find(&user.Wishlist)
	//respond with the wishlist
	c.JSON(http.StatusOK, gin.H{
		"Wishlist": user.Wishlist,
	})

}

func AddToWishlist(c *gin.Context) {
	DB := initializers.DB

	var body struct {
		CustomerID  uint
		ProductID   uint
		Name        string
		Description string
		Available   bool
		Price       int
	}
	//bind the body with the context
	if c.Bind(&body) != nil {
		helper.ShowError(c, "failed to read the body")
		return
	}
	// check if customer exists
	var user models.User
	profile := DB.First(&user, body.CustomerID)

	if profile.Error != nil {
		helper.ShowError(c, "Customer is not found")
		return
	}

	//check product already exists or not
	var wishlist models.WishlistOfUser
	result := DB.Model(&wishlist).Where("product_number = ? AND wishlist_belongsto = ?", body.ProductID, body.CustomerID).Find(&wishlist)
	if result.RowsAffected > 0 {
		helper.ShowError(c, "product is already there")
		return
	}

	//Add product in Wishlist
	DB.Model(&user).Association("Wishlist").Append(&models.WishlistOfUser{
		ProductNumber: body.ProductID,
		Name:          body.Name,
		Description:   body.Description,
		Available:     body.Available,
		Price:         body.Price})

	c.JSON(http.StatusOK, gin.H{
		"msg": "Added product in the wishlist successfully",
	})
}

func Removefromwishlist(c *gin.Context) {
	DB := initializers.DB

	var body struct {
		CustomerID uint
		ProductID  uint
	}
	//bind the body with the context
	if c.Bind(&body) != nil {
		helper.ShowError(c, "failed to read the body")
		return
	}
	// check if customer exists
	var user models.User
	profile := DB.First(&user, body.CustomerID)

	if profile.Error != nil {
		helper.ShowError(c, "Customer is not found")
		return
	}

	//check product already exists or not
	//Delete the product from the Wishlist

	var wishlist models.WishlistOfUser
	DB.Model(&wishlist).Where("product_number = ? AND wishlist_belongsto = ?", body.ProductID, body.CustomerID).Delete(&wishlist)
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully deleted message",
	})

}
