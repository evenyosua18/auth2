package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MRegistrationUsecase struct {
	mock.Mock
}

func (m *MRegistrationUsecase) RegistrationUser(ctx context.Context, in interface{}) (interface{}, error) {
	args := m.Called(ctx, in)

	return args.Get(0), args.Error(1)
}
