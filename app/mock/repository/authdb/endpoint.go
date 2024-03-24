package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MEndpointRepository struct {
	mock.Mock
}

func (m *MEndpointRepository) InsertEndpoint(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MEndpointRepository) UpdateEndpoint(ctx context.Context, filter, in interface{}) error {
	args := m.Called(ctx, in, filter)

	return args.Error(0)
}

func (m *MEndpointRepository) GetEndpoint(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}

func (m *MEndpointRepository) GetEndpoints(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}
