package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MAccessTokenRepository struct {
	mock.Mock
}

func (m *MAccessTokenRepository) InsertAccessToken(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MAccessTokenRepository) GetAccessToken(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}

func (m *MAccessTokenRepository) DeleteAccessToken(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}

func (m *MAccessTokenRepository) UpdateAccessToken(ctx context.Context, filter, in interface{}) error {
	args := m.Called(ctx, filter, in)
	return args.Error(0)
}
