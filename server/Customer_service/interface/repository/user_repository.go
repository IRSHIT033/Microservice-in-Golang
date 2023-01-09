package repository

import (
	"errors"

	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/domain/model"
	"github.com/IRSHIT033/E-comm-GO-/server/Customer_service/usecase/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Find(user *model.User) (*model.User, error) {

	err := ur.db.First(&user, "email= ?", user.Email).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (ur *userRepository) Save(user *model.User) (*model.User, error) {

	err := ur.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) AddProduct(product *model.Product) (string, error) {
	var msg string
	// check if customer exists

	var user model.User

	err := ur.db.First(&user, product.AddedBy).Error

	if err != nil {
		return " Customer is not found ", err
	}

	//check product already exists or not

	result := ur.db.Model(&product).Where("product_id = ? AND added_by = ?", product.ProductID, product.AddedBy).Find(&product)
	if result.RowsAffected > 0 {
		return " product already exists ", err
	}

	//Add product in Cart

	ur.db.Model(&user).Association("Cart").Append(&model.Product{
		ProductID:       product.ProductID,
		ProductImageSrc: product.ProductImageSrc,
		Name:            product.Name,
		Price:           product.Price,
		Unit:            product.Unit,
	})

	msg = "Added product in the cart successfully"
	return msg, nil
}

func (ur *userRepository) GetCart(userId uint) ([]*model.Product, error) {
	var products []*model.Product
	var user model.User
	err := ur.db.First(&user, userId).Error

	if err != nil {
		return nil, errors.New("can't find customers")
	}

	//get the cart of a specific user
	ur.db.Model(&user).Association("Cart").Find(&products)

	return products, nil
}

func (ur *userRepository) RemoveProduct(userId uint, productId uint) (string, error) {
	var msg string
	// check if customer exists
	var user model.User
	err := ur.db.First(&user, userId).Error

	if err != nil {

		return "customer not found", err
	}

	//Delete the product from the CART

	var cart model.Product
	ur.db.Model(&cart).Where("product_id = ? AND added_by = ?", productId, userId).Delete(&cart)

	msg = " product removed from the cart successfully"
	return msg, nil
}
