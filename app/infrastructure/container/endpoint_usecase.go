//go:build wireinject
// +build wireinject

package container

import (
	endpointRepo "github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	"github.com/evenyosua18/auth2/app/usecase/endpoint"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeEndpointUsecase(connection *db.MongoConnection) (endpointUC endpoint.IRegistrationEndpointUsecase) {
	wire.Build(endpointRepo.NewEndpointRepository, endpoint.NewUsecaseRegistrationEndpoint)
	return
}
