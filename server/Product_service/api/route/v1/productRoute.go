package route

import (
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/api/controller"
	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/repository_product"
	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	pr := repository_product.NewProductRepository(db)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
	}

	group.POST("/createProduct", pc.CreateProductData)
	group.GET("/fetchallProduct", pc.FetchAllProducts)
	group.GET("/fetchbycategory", pc.FetchbyCategory)
	group.POST("/fetch/:id", pc.FetchbyId)

}
