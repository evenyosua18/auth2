package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/tracing"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
)

func (s *ServiceScope) CreateScope(ctx context.Context, in *pb.CreateScopeRequest) (*empty.Empty, error) {
	// tracing
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "GetScope"))
	defer tracing.Close(sp)

	// create scope
	if err := s.uc.CreateScope(tracing.Context(sp), in); err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	return &empty.Empty{}, nil
}
