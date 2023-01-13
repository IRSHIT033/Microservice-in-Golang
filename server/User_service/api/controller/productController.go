package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain_user.ProductUseCase
}

func (pc *ProductController) AddProductToCustomers(c *gin.Context) {
	var product domain_user.Product

	err := c.ShouldBind(&product)

	if err != nil {
		fmt.Print("1")
		c.JSON(http.StatusBadRequest, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	fmt.Println(userID)
	userid, err := strconv.ParseUint(userID, 10, 32)
	product.AddedBy = uint(userid)

	fmt.Println(product.AddedBy)

	if err != nil {
		fmt.Print("2")
		c.JSON(http.StatusBadRequest, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	err = pc.ProductUsecase.Add(c, &product)
	if err != nil {
		fmt.Print("3")
		c.JSON(http.StatusInternalServerError, domain_user.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain_user.SuccessResponse{
		Message: "product added to cart successfully",
	})

}
