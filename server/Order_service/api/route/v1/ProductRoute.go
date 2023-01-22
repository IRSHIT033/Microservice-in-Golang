package route

import (
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/api/controller"
	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/repository_order"
	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewOrderRoute(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	pr := repository_order.NewOrderRepository(db)
	pc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(pr, timeout),
	}
	group.POST("/create", pc.CreateOrder)
	group.POST("/fetch", pc.FetchOrders)
}
