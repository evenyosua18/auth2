package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) DeleteScope(ctx context.Context, filter interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// delete scope
	if res, err := r.col.DeleteOne(ctx, f.Filter()); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, res)
		return nil
	}
}
