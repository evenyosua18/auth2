package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) ListScope(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.ScopeFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// get scopes
	cur, err := r.col.Find(ctx, f.Filter(), f.Pagination.Option(), f.Sort.Option())

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	defer cur.Close(ctx)

	// loop for mapping model
	var scopes []model.ScopeModel
	for cur.Next(ctx) {
		// decode
		var scope model.ScopeModel
		if err := cur.Decode(&scope); err != nil {
			return nil, tracing.LogError(sp, codes.Wrap(err, 502))
		} else {
			scopes = append(scopes, scope)
		}
	}

	tracing.LogResponse(sp, scopes)
	return scopes, nil
}
