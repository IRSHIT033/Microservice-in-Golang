package router

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, con controller.AppController) *gin.Engine {

	router.POST("/login", func(context *gin.Context) { con.User.GetUser(context) })
	router.POST("/signup", func(context *gin.Context) { con.User.CreateUser(context) })
	router.PUT("/addtoCart", func(context *gin.Context) { con.User.AddToCart(context) })
	router.GET("/getCart/:id", func(context *gin.Context) { con.User.GetCart(context) })
	router.DELETE("/removefromCart", func(context *gin.Context) { con.User.RemoveFromCart(context) })

	return router

}
