package authdb

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/model"
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
	res := args.Get(0)

	if res == nil {
		return res, args.Error(1)
	} else if _, ok := res.(model.RefreshTokenModel); !ok {
		return nil, errors.New("invalid object")
	} else if ok {
		return res, nil
	}

	return nil, errors.New("invalid args")
}

func (m *MRefreshTokenRepository) DeleteRefreshToken(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}
