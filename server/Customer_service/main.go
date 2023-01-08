package main

import (
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/database"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/registry"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.CreatedatabaseInstance()
	reg := registry.NewRegistry(db)
	routes := gin.Default()
	routes = router.NewRouter(routes, reg.NewAppController())
	//port := os.Getenv("PORT")
	routes.Run(":5000")
}
