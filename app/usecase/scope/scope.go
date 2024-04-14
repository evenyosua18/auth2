package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	"github.com/evenyosua18/auth2/app/repository/authdb/scope"
)

type IScopeUsecase interface {
	CreateScope(ctx context.Context, in interface{}) error
	UpdateScope(ctx context.Context, in, filter interface{}) error
	DeleteScope(ctx context.Context, filter interface{}) error
	GetScope(ctx context.Context, filter interface{}) (interface{}, error)
	GetScopes(ctx context.Context, filter interface{}) (interface{}, interface{}, error)
}

type UsecaseScope struct {
	scope scope.IScopeRepository

	endpoint endpoint.IEndpointRepository
}

func NewScopeUsecase(scopeRepository scope.IScopeRepository, endpointRepository endpoint.IEndpointRepository) IScopeUsecase {
	return &UsecaseScope{scope: scopeRepository, endpoint: endpointRepository}
}
