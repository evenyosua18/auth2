package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryAccessToken) GetAccessToken(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// filter
	var f model.AccessTokenFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// call database
	var res *model.AccessTokenModel
	if err := r.col.FindOne(ctx, f.Filter()).Decode(&res); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
