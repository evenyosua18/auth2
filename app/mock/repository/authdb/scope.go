package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MScopeRepository struct {
	mock.Mock
}

func (m *MScopeRepository) InsertScope(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MScopeRepository) GetScope(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}

func (m *MScopeRepository) DeleteScope(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}

func (m *MScopeRepository) UpdateScope(ctx context.Context, in, filter interface{}) error {
	args := m.Called(ctx, in, filter)

	return args.Error(0)
}

func (m *MScopeRepository) ListScope(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}

func (m *MScopeRepository) CountScope(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}
