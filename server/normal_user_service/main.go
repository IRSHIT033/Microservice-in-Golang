package main

import (
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/controllers"
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/initializers"
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Envinitializer()
	initializers.Connect_DB()
	initializers.Database_sync()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Authorization, controllers.Validate)

	r.Run()
}
