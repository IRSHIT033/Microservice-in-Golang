package initializers

import (
	"github.com/IRSHIT033/E-comm-GO-/server/normal_user_service/models"
)

func Database_sync() {
	DB.AutoMigrate(
		&models.Address{},
		&models.WishlistOfUser{},
		&models.Order{},
		&models.Product{}, &models.User{})
}
