package controller

import (
	"fmt"
	"net/http"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/domain_order"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderUsecase domain_order.OrderUsecase
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var request domain_order.CreateOrderRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain_order.ErrorResponse{Message: err.Error()})
		return
	}

	err = oc.OrderUsecase.Create(c, request.CustomerId, request.TransactionId)
	if err != nil {
		c.JSON(http.StatusNotFound, domain_order.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain_order.SuccessResponse{
		Message: "order saved successfully",
	})
}

func (oc *OrderController) FetchOrders(c *gin.Context) {
	var request domain_order.FetchOrderRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain_order.ErrorResponse{Message: err.Error()})
		return
	}
	fmt.Println(request.CustomerId)
	all_orders, err := oc.OrderUsecase.Fetch(c, request.CustomerId)
	if err != nil {
		c.JSON(http.StatusNotFound, domain_order.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, all_orders)

}
