package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (s *ServiceScope) GetScopes(ctx context.Context, filter *pb.GetScopesRequest) (*pb.GetScopesResponse, error) {
	// tracing
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "GetScopes"))
	defer tracing.Close(sp)

	// get filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// set sort
	f.Sort.SetDefault(nil)

	// get scopes
	scopes, pagination, err := s.uc.GetScopes(ctx, f)

	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	// update pagination data
	f.Pagination = pagination.(model.Pagination)
	tracing.LogObject(sp, "filter", f)

	tracing.LogResponse(sp, scopes)
	return s.out.GetScopesResponse(tracing.Context(sp), scopes, f)
}
