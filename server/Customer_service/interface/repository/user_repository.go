package repository

import (
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

func (ur *userRepository) Find(u *model.User) (*model.User, error) {

	var user *model.User

	err := ur.db.First(&user, "email= ?", u.Email).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}
