package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
)

func (s *ServiceAccessToken) ValidateAccessToken(ctx context.Context, request *pb.ValidateTokenRequest) (*empty.Empty, error) {
	// tracing
	sp := tracing.StartChild(ctx, request)
	defer tracing.Close(sp)

	// call usecase
	if err := s.uc.ValidateAccessToken(tracing.Context(sp), request); err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	return &empty.Empty{}, nil
}
