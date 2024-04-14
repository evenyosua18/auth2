package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/usecase/scope"
)

type IScopeBuilder interface {
	GetScopeResponse(ctx context.Context, in interface{}) (*pb.GetScopeResponse, error)
	GetScopesResponse(ctx context.Context, result, pagination interface{}) (*pb.GetScopesResponse, error)
}

type ServiceScope struct {
	uc  scope.IScopeUsecase
	out IScopeBuilder
	pb.UnimplementedScopeServiceServer
}

func NewScopeService(out IScopeBuilder, uc scope.IScopeUsecase) *ServiceScope {
	return &ServiceScope{
		uc:  uc,
		out: out,
	}
}
