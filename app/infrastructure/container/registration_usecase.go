//go:build wireinject
// +build wireinject

package container

import (
	"github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	endpointRepo "github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	"github.com/evenyosua18/auth2/app/usecase/endpoint"
	register "github.com/evenyosua18/auth2/app/usecase/registration"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeRegistrationUserUsecase(connection *db.MongoConnection) (registration register.IRegistrationUserUsecase) {
	wire.Build(user.NewUserRepository, accesstoken.NewAccessTokenRepository, refreshtoken.NewRefreshTokenRepository, register.NewRegistrationUserUsecase)
	return
}

func InitializeRegistrationEndpointUsecase(connection *db.MongoConnection) (endpointUC endpoint.IRegistrationEndpointUsecase) {
	wire.Build(endpointRepo.NewEndpointRepository, endpoint.NewUsecaseRegistrationEndpoint)
	return
}
