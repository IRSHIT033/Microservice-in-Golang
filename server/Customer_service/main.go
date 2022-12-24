package main

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/controllers"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	//Env variable initializer
	initializers.Envinitializer()
	//data base connection
	initializers.Connect_DB()
	initializers.Database_sync()
}

func main() {
	r := gin.Default()
	//create user
	r.POST("/signup", controllers.SignUp)
	//login user
	r.POST("/login", controllers.Login)
	//user authenticator
	r.GET("/validate", middleware.Authorization, controllers.Validate)

	//get wishlist of a user
	r.GET("/wishlist/:id", middleware.Authorization, controllers.GetWishlist)
	//add item into wishlist of a user
	r.PATCH("/addtowishlist", middleware.Authorization, controllers.AddToWishlist)
	//remove item from wishlist of a user
	r.PATCH("/removefromwishlist", middleware.Authorization, controllers.Removefromwishlist)

	//get cart of a user
	r.GET("/getCart/:id", middleware.Authorization, controllers.GetCart)
	//add item to cart
	r.PATCH("/addtocart", middleware.Authorization, controllers.AddToCart)
	//remove item from cart
	r.PATCH("/removefromCart", middleware.Authorization, controllers.RemoveFromCart)

	r.Run()
}
