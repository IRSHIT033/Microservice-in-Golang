package usecase

import (
	"context"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domain_user.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain_user.UserRepository, timeout time.Duration) domain_user.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) CreateUser(c context.Context, user *domain_user.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain_user.User, int, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain_user.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain_user.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)

}
