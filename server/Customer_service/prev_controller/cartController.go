package controller

import (
	"net/http"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/helper"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/models"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {

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

	var cart []models.Product

	//get the cart of a specific user
	DB.Model(&user).Association("Cart").Find(&cart)

	//respond with cart
	c.JSON(http.StatusOK, gin.H{
		"result": cart,
	})

}

func AddToCart(c *gin.Context) {

	DB := initializers.DB

	var body struct {
		CustomerID      uint
		ProductID       uint
		ProductImageSrc string
		Name            string
		Price           int
		Unit            int
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
	var cart models.Product
	result := DB.Model(&cart).Where("product_id = ? AND added_by = ?", body.ProductID, body.CustomerID).Find(&cart)
	if result.RowsAffected > 0 {
		helper.ShowError(c, "product is already there")
		return
	}

	//Add product in Cart
	DB.Model(&user).Association("Cart").Append(&models.Product{
		ProductID:       body.ProductID,
		ProductImageSrc: body.ProductImageSrc,
		Name:            body.Name,
		Price:           body.Price,
		Unit:            body.Unit,
	})

	//respond with success msg

	c.JSON(http.StatusOK, gin.H{
		"msg": "Added product in the wishlist successfully",
	})
}

func RemoveFromCart(c *gin.Context) {
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

	//Delete the product from the CART

	var cart models.Product
	DB.Model(&cart).Where("product_id = ? AND added_by = ?", body.ProductID, body.CustomerID).Delete(&cart)
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully removed item from cart",
	})

}
