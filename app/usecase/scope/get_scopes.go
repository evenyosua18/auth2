package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (u *UsecaseScope) GetScopes(ctx context.Context, filter interface{}) (interface{}, interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// if pagination, get total data
	if f.Pagination.CurrentPage != 0 {
		// get total data
		total, err := u.scope.CountScope(tracing.Context(sp), filter)

		if err != nil {
			return nil, nil, tracing.LogError(sp, err)
		}

		// update pagination data
		f.Pagination.AddTotalPage(total)

		tracing.LogObject(sp, "filter", f)
	}

	// get scopes
	scopes, err := u.scope.ListScope(tracing.Context(sp), f)

	if err != nil {
		return nil, nil, tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, scopes)
	return scopes, f.Pagination, nil
}
