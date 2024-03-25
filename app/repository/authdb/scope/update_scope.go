package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) UpdateScope(ctx context.Context, in, filter interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// decode model
	var scope model.ScopeModel
	if err := mapstructure.Decode(in, &scope); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// update
	if res, err := r.col.UpdateOne(ctx, f.Filter(), scope.Update()); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, res)
		return nil
	}
}
