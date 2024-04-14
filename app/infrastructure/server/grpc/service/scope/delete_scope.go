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

func (s *ServiceScope) DeleteScope(ctx context.Context, scopeRequest *pb.DeleteScopeRequest) (*empty.Empty, error) {
	// tracing
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "DeleteScope"))
	defer tracing.Close(sp)

	if err := s.uc.DeleteScope(tracing.Context(sp), scopeRequest); err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	return &empty.Empty{}, nil
}
