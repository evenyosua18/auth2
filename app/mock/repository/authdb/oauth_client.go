package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MOauthClientRepository struct {
	mock.Mock
}

func (m *MOauthClientRepository) InsertOauthClient(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MOauthClientRepository) GetOauthClient(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}
