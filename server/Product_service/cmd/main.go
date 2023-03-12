package main

import (
	"time"

	routeV1 "github.com/IRSHIT033/E-comm-GO-/server/Product_service/api/route/v1"
	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/bootstrap"
	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/grpc_config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	db := app.DB
	gin := gin.Default()
	productRoutes := gin.Group("product")

	timeout := time.Duration(24) * time.Second
	routeV1.Setup(db, timeout, productRoutes)

	go grpc_config.GRPCListen(db)
	gin.Run(":3001")

}
