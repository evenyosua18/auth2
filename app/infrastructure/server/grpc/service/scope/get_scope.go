package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/tracing"
	"google.golang.org/grpc/status"
)

func (s *ServiceScope) GetScope(ctx context.Context, in *pb.GetScopeRequest) (*pb.GetScopeResponse, error) {
	// trace
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "GetScope"))
	defer tracing.Close(sp)

	// get scope
	res, err := s.uc.GetScope(tracing.Context(sp), in)

	if err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	tracing.LogResponse(sp, res)
	return s.out.GetScopeResponse(ctx, res)
}
