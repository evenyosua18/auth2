package registration

import (
	"context"
	"github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
)

type IRegistrationUsecase interface {
	RegistrationUser(ctx context.Context, in interface{}) (interface{}, error)
}

type UsecaseRegistration struct {
	user         user.IUserRepository
	accessToken  accesstoken.IAccessTokenRepository
	refreshToken refreshtoken.IRefreshTokenRepository
}

func NewRegistrationUsecase(userRepo user.IUserRepository, accessTokenRepo accesstoken.IAccessTokenRepository, refreshTokenRepo refreshtoken.IRefreshTokenRepository) IRegistrationUsecase {
	return &UsecaseRegistration{
		user:         userRepo,
		accessToken:  accessTokenRepo,
		refreshToken: refreshTokenRepo,
	}
}
