package authdb

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/model"
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
	res := args.Get(0)

	if res == nil {
		return res, args.Error(1)
	} else if _, ok := res.(*model.AccessTokenModel); !ok {
		return nil, errors.New("invalid object")
	} else if ok {
		return res, nil
	}

	return nil, errors.New("invalid args")
}

func (m *MAccessTokenRepository) DeleteAccessToken(ctx context.Context, filter interface{}) error {
	args := m.Called(ctx, filter)

	return args.Error(0)
}

func (m *MAccessTokenRepository) UpdateAccessToken(ctx context.Context, filter, in interface{}) error {
	args := m.Called(ctx, filter, in)
	return args.Error(0)
}
