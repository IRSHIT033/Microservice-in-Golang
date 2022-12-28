package initializers

import (
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/models"
)

func Database_sync() {
	DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.WishlistOfUser{},
		&models.Order{},
		&models.Product{})
}
