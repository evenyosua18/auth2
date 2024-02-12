package refreshtoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryRefreshToken) GetRefreshToken(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// get filter
	var f model.RefreshTokenFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	cur, err := r.col.Aggregate(ctx, f.Aggregate())

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	// decode result
	var refreshToken model.RefreshTokenModel
	if cur.Next(ctx) {
		if err := cur.Decode(&refreshToken); err != nil {
			return nil, tracing.LogError(sp, codes.Wrap(err, 501))
		}
	}

	tracing.LogResponse(sp, refreshToken)
	return refreshToken, nil
}
