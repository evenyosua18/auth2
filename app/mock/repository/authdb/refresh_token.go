package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MRefreshTokenRepository struct {
	mock.Mock
}

func (m *MRefreshTokenRepository) InsertRefreshToken(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MRefreshTokenRepository) GetRefreshToken(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)
	return args.Get(0), args.Error(1)
}

func (m *MRefreshTokenRepository) DeleteRefreshToken(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}
