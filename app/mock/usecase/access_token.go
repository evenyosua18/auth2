package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MAccessTokenUsecase struct {
	mock.Mock
}

func (m *MAccessTokenUsecase) PasswordGrant(ctx context.Context, in interface{}) (interface{}, error) {
	args := m.Called(ctx, in)

	return args.Get(0), args.Error(1)
}

func (m *MAccessTokenUsecase) ValidateAccessToken(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MAccessTokenUsecase) RefreshAccessToken(ctx context.Context, in interface{}) (interface{}, error) {
	args := m.Called(ctx, in)

	return args.Get(0), args.Error(1)
}
