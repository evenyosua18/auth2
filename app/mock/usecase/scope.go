package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MScopeUsecase struct {
	mock.Mock
}

func (m *MScopeUsecase) CreateScope(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MScopeUsecase) UpdateScope(ctx context.Context, in, filter interface{}) error {
	args := m.Called(ctx, in, filter)

	return args.Error(0)
}

func (m *MScopeUsecase) DeleteScope(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}

func (m *MScopeUsecase) GetScope(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}

func (m *MScopeUsecase) GetScopes(ctx context.Context, filter interface{}) (interface{}, interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Get(1), args.Error(2)
}
