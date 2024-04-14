package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) CountScope(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// count
	total, err := r.col.CountDocuments(ctx, f.Filter())

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, total)
	return total, nil
}
