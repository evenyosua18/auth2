package user

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryUser) GetUser(ctx context.Context, filter interface{}) (interface{}, error) {
	// start tracer
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// filter
	var f model.UserFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// call database
	var res *model.UserModel
	if err := r.col.FindOne(ctx, f.Filter()).Decode(&res); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
