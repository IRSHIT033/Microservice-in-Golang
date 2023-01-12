package repository_user

import (
	"context"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain_user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain_user.User) error {
	err := ur.db.Create(&user).Error
	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain_user.User, int, error) {
	var user domain_user.User
	result := ur.db.Where("email = ?", email).Find(&user)
	userExists := int(result.RowsAffected)
	err := result.Error
	return user, userExists, err
}
