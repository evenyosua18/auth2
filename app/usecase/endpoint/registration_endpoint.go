package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/repository/authdb/endpoint"
	"google.golang.org/grpc"
)

type IRegistrationEndpointUsecase interface {
	RegisterGRPC(context.Context, map[string]grpc.ServiceInfo) interface{}
}

type UsecaseRegistrationEndpoint struct {
	endpoint       endpoint.IEndpointRepository
	IgnorePrefixes map[string]bool
}

func NewUsecaseRegistrationEndpoint(endpointUC endpoint.IEndpointRepository) IRegistrationEndpointUsecase {
	// ignore prefix
	ignorePrefixes := map[string]bool{
		"grpc": true,
	}

	return &UsecaseRegistrationEndpoint{endpoint: endpointUC, IgnorePrefixes: ignorePrefixes}
}
