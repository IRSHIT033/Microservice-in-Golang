package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, timeout time.Duration, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	NewSignupRouter(timeout, db, publicRouterV1)
	NewLoginRouter(timeout, db, publicRouterV1)
}
