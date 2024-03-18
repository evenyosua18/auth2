package authdb

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MUserRepository struct {
	mock.Mock
}

func (m *MUserRepository) InsertUser(ctx context.Context, in interface{}) error {
	args := m.Called(ctx, in)

	return args.Error(0)
}

func (m *MUserRepository) GetUser(ctx context.Context, filter interface{}) (interface{}, error) {
	args := m.Called(ctx, filter)

	return args.Get(0), args.Error(1)
}
