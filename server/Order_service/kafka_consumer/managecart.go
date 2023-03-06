package kafka_consumer

import (
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/domain_order"
	"gorm.io/gorm"
)

func AddToCart(db *gorm.DB, customerId uint, product domain_order.Product) error {

	cart := domain_order.Cart{CustomerId: customerId}
	log.Println(customerId)
	result := db.Model(&cart).Where("customer_id = ?", customerId).Find(&cart)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		db.Create(&cart)
	}

	db.Model(&cart).Where("customer_id = ?", customerId).Association("Products").Append(&product)
	log.Println("added successfully......")

	return nil

}

func RemoveToCart(db *gorm.DB, customerId uint, product domain_order.Product) error {
	cart := domain_order.Cart{CustomerId: customerId}
	result := db.Model(&cart).Where("customer_id = ?", customerId).Find(&cart)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		db.Create(&cart)
	}

	db.Model(&cart).Association("Products").Delete(&product)
	log.Println("added successfully......")
	return nil
}
