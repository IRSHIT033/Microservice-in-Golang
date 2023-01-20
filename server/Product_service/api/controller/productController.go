package controller

import (
	"net/http"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/domain_product"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain_product.ProductUsecase
}

func (pc *ProductController) CreateProductData(c *gin.Context) {
	var product domain_product.Product
	err := c.ShouldBind(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	err = pc.ProductUsecase.Create(c, product)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain_product.SuccessResponse{
		Message: "product saved successfully",
	})

}

func (pc *ProductController) FetchAllProducts(c *gin.Context) {

	products, err := pc.ProductUsecase.Fetch(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) FetchbyId(c *gin.Context) {

	productid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	productID := uint(productid)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	product, err := pc.ProductUsecase.FetchbyId(c, productID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (pc *ProductController) FetchbyCategory(c *gin.Context) {
	body := domain_product.CategoryRequest{}

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	products, err := pc.ProductUsecase.FetchbyCategory(c, body.Category)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain_product.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)

}
