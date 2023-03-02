package route

import (
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/api/controller"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/repository_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {

	pr := repository_user.NewProductRepository(db)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
	}

	group.PUT("/addproductToCart", pc.AddProductToCustomers)
	group.PUT("/removeFromCart/:id", pc.RemoveProductFromCart)
	group.GET("/getcartwithUser", pc.GetCartWithUser)
	group.GET("/getCartofauser", pc.GetproductOfUser)

}
