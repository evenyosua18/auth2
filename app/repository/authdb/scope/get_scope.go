package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) GetScope(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	tracing.LogObject(sp, "decode filter", f.Filter())

	// get scope
	var scope model.ScopeModel
	if err := r.col.FindOne(ctx, f.Filter()).Decode(&scope); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, scope)
		return scope, nil
	}
}
