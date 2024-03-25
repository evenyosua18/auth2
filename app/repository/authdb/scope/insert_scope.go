package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryScope) InsertScope(ctx context.Context, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx)
	defer tracing.Close(sp)

	// set model
	scope := model.ScopeModel{}.SetCreate()
	if err := mapstructure.Decode(in, &scope); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	if res, err := r.col.InsertOne(ctx, scope); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, res)
		return nil
	}
}
