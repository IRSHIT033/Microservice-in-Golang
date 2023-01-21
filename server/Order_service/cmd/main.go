package main

import (
	"time"

	routeV1 "github.com/IRSHIT033/E-comm-GO-/server/Order_service/api/route/v1"
	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	db := app.DB
	gin := gin.Default()
	routerV1 := gin.Group("order")
	timeout := time.Duration(24) * time.Second

	routeV1.Setup(db, timeout, routerV1)

	gin.Run(":5003")
}
