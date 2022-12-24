package controllers

import (
	"net/http"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/models"
	"github.com/gin-gonic/gin"
)

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
		"msg": "successfully removed item from wishlist",
	})

}
