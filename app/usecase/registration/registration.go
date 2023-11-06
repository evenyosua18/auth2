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

type IRegistrationBuilder interface {
	RegistrationUserResponse(ctx context.Context, in interface{}, code int) (interface{}, error)
}

type UsecaseRegistration struct {
	out          IRegistrationBuilder
	user         user.IUserRepository
	accessToken  accesstoken.IAccessTokenRepository
	refreshToken refreshtoken.IRefreshTokenRepository
}

func NewRegistrationUsecase(out IRegistrationBuilder, userRepo user.IUserRepository, accessTokenRepo accesstoken.IAccessTokenRepository, refreshTokenRepo refreshtoken.IRefreshTokenRepository) IRegistrationUsecase {
	return &UsecaseRegistration{
		out:          out,
		user:         userRepo,
		accessToken:  accessTokenRepo,
		refreshToken: refreshTokenRepo,
	}
}
