package builder

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/stretchr/testify/mock"
)

type MAccessTokenBuilder struct {
	mock.Mock
}

func (m *MAccessTokenBuilder) AccessTokenResponse(ctx context.Context, in interface{}, code int) (*pb.AccessTokenResponse, error) {
	args := m.Called(ctx, in, code)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else if res, ok := args.Get(0).(*pb.AccessTokenResponse); !ok {
		return nil, errors.New("invalid testing object")
	} else {
		return res, args.Error(1)
	}
}
