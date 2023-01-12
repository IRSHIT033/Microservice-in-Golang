package bootstrap

import (
	"fmt"
	"os"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// this function returns a database intstance
func CreatedatabaseInstance() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to conenct to db")
	}
	//migrate all databases

	db.AutoMigrate(
		&domain_user.User{},
		&domain_user.Product{},
		&domain_user.Address{},
	)

	return db
}
