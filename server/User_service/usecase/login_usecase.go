package usecase

import (
	"context"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain_user.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain_user.UserRepository, timeout time.Duration) domain_user.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain_user.User, int, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain_user.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain_user.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
