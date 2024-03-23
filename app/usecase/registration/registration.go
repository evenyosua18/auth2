package registration

import (
	"context"
	"github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	"google.golang.org/grpc"
)

/*Registration User*/

type IRegistrationUserUsecase interface {
	RegistrationUser(ctx context.Context, in interface{}) (interface{}, error)
}

type UsecaseRegistrationUser struct {
	user         user.IUserRepository
	accessToken  accesstoken.IAccessTokenRepository
	refreshToken refreshtoken.IRefreshTokenRepository
}

func NewRegistrationUserUsecase(userRepo user.IUserRepository, accessTokenRepo accesstoken.IAccessTokenRepository, refreshTokenRepo refreshtoken.IRefreshTokenRepository) IRegistrationUserUsecase {
	return &UsecaseRegistrationUser{
		user:         userRepo,
		accessToken:  accessTokenRepo,
		refreshToken: refreshTokenRepo,
	}
}

/*Registration Endpoint*/

type IRegistrationEndpointUsecase interface {
	RegisterGRPC(context.Context, map[string]grpc.ServiceInfo) interface{}
}

type UsecaseRegistrationEndpoint struct {
	endpoint       endpoint.IEndpointRepository
	IgnorePrefixes map[string]bool
}

func NewRegistrationEndpointUsecase(endpointUC endpoint.IEndpointRepository) IRegistrationEndpointUsecase {
	// ignore prefix
	ignorePrefixes := map[string]bool{
		"grpc": true,
	}

	return &UsecaseRegistrationEndpoint{endpoint: endpointUC, IgnorePrefixes: ignorePrefixes}
}
