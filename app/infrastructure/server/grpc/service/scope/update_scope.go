package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/golang/protobuf/ptypes/empty"
)

type filter struct {
	Id string
}

type request struct {
	Name      string
	Service   string
	Endpoints []string
}

func (s *ServiceScope) UpdateScope(ctx context.Context, in *pb.UpdateScopeRequest) (*empty.Empty, error) {
	// tracing
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "UpdateScope"))
	defer tracing.Close(sp)

	// update scope
	if err := s.uc.UpdateScope(ctx, request{
		Name:      in.Name,
		Service:   in.Service,
		Endpoints: in.Endpoints,
	}, filter{
		Id: in.Id,
	}); err != nil {
		return &empty.Empty{}, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	return &empty.Empty{}, nil
}
