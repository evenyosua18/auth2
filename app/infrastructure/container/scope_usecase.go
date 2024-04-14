//go:build wireinject
// +build wireinject

package container

import (
	"github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	scopeRepo "github.com/evenyosua18/auth2/app/repository/authdb/scope"
	"github.com/evenyosua18/auth2/app/usecase/scope"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeScopeUsecase(connection *db.MongoConnection) (scopeUC scope.IScopeUsecase) {
	wire.Build(scopeRepo.NewScopeRepository, endpoint.NewEndpointRepository, scope.NewScopeUsecase)
	return
}
