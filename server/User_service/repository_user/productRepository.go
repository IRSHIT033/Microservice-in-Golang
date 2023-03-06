package repository_user

import (
	"context"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/kafka_producer"
	"gorm.io/gorm"
)

type productRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain_user.Productrepository {
	return &productRepository{
		database: db,
	}
}

func (pr *productRepository) Add(c context.Context, customerId uint, product *domain_user.Product) error {
	// find user
	var user domain_user.User
	err := pr.database.First(&user, customerId).Error

	if err != nil {
		return err
	}
	// // check if product already exists in user's cart or not
	// result := pr.database.Model(&product).Where("product_id = ? AND added_by = ?", product.ProductID, product.AddedBy).Find(&product)

	// if result.Error != nil {
	// 	return result.Error
	// }

	// if result.RowsAffected > 0 {
	// 	return fmt.Errorf(" product already exist in cart")
	// }

	// add product to the user's cart
	pr.database.Model(&user).Association("Cart").Append(&domain_user.Product{
		Model:           gorm.Model{ID: product.ID},
		ProductImageSrc: product.ProductImageSrc,
		Name:            product.Name,
		Price:           product.Price,
		Unit:            product.Unit,
	})

	return nil

}

func (pr *productRepository) FetchByUserID(c context.Context, userID uint) ([]domain_user.Product, error) {
	var user domain_user.User
	var cart []domain_user.Product

	//Check if the customer exists
	err := pr.database.First(&user, userID).Error

	if err != nil {
		return cart, err
	}

	// get cart of user
	pr.database.Model(&user).Association("Cart").Find(&cart)

	return cart, nil
}

func (pr *productRepository) Remove(c context.Context, productID uint, userID uint) error {
	var product domain_user.Product
	var user domain_user.User
	// find user

	err := pr.database.First(&user, userID).Error
	if err != nil {
		return err
	}

	err = pr.database.First(&product, productID).Error
	if err != nil {
		return err
	}

	//send the "adding product to cart" message to kafka broker
	var cart domain_user.KafkaMessagePayload
	cart.CustomerId = userID
	cart.Product = product
	cart.Operation = "Remove"
	go kafka_producer.ProduceCart(cart)
	////////////////////////////////////////////////////////////

	pr.database.Model(&user).Association("Cart").Delete(&product)

	return nil
}

func (pr *productRepository) FetchUserCart(c context.Context, userID uint) (domain_user.User, error) {
	var user domain_user.User

	//Check if the customer exists
	err := pr.database.First(&user, userID).Error

	if err != nil {
		return user, err
	}

	// get cart of user
	pr.database.Preload("Cart").Where("id = ?", userID).Find(&user)

	return user, nil
}
