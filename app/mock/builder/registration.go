package builder

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/stretchr/testify/mock"
)

type MRegistrationBuilder struct {
	mock.Mock
}

func (m *MRegistrationBuilder) RegistrationUserResponse(ctx context.Context, in interface{}, code int) (*pb.RegistrationUserResponse, error) {
	args := m.Called(ctx, in, code)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else if res, ok := args.Get(0).(*pb.RegistrationUserResponse); !ok {
		return nil, errors.New("invalid testing object")
	} else {
		return res, args.Error(1)
	}
}
