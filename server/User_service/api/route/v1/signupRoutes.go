package route

import (
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/api/controller"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/repository_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSignupRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository_user.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	}
	group.POST("/signup", sc.Signup)

}
