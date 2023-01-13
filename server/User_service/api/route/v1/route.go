package route

import (
	"os"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, timeout time.Duration, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	NewSignupRouter(timeout, db, publicRouterV1)
	NewLoginRouter(timeout, db, publicRouterV1)

	protectedRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	secret_key := os.Getenv("SECRET_KEY")
	protectedRouterV1.Use(middleware.JwtAuthMiddleware(secret_key))
	NewProductRoute(timeout, db, protectedRouterV1)
}
