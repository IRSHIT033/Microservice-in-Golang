package repository_user

import (
	"context"
	"fmt"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
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

func (pr *productRepository) Add(c context.Context, product *domain_user.Product) error {
	//find user
	var user domain_user.User
	err := pr.database.First(&user, product.AddedBy).Error

	if err != nil {
		return err
	}
	// check if product already exists in user's cart or not
	result := pr.database.Model(&product).Where("product_id = ? AND added_by = ?", product.ProductID, product.AddedBy).Find(&product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return fmt.Errorf(" product already exist in cart")
	}

	// add product to the user's cart
	pr.database.Model(&user).Association("Cart").Append(&domain_user.Product{
		ProductID:       product.ProductID,
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

	//get cart of user
	pr.database.Model(&user).Association("Cart").Find(&cart)

	return cart, nil
}

func (pr *productRepository) Remove(c context.Context, productID uint, userID uint) error {
	var product domain_user.Product
	var user domain_user.User
	//find user

	err := pr.database.First(&user, userID).Error

	if err != nil {
		return err
	}

	pr.database.Model(&product).Where("product_id = ? AND added_by = ?", productID, userID).Delete(&product)

	return nil
}
