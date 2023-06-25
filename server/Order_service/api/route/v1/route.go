package route

import (
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, timeout time.Duration, routerV1 *gin.RouterGroup) {
	protectedRouterV1 := routerV1.Group("")
	protectedRouterV1.Use(middleware.AuthMiddleware())
	NewOrderRoute(timeout, db, protectedRouterV1)
}
