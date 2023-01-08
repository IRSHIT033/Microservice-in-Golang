package router

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, con controller.AppController) *gin.Engine {
	router.POST("/login", func(context *gin.Context) { con.User.GetUser(context) })

	return router
}
