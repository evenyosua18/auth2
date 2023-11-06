package refreshtoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepositoryRefreshToken) GetRefreshToken(ctx context.Context, agg []bson.M) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, agg)
	defer tracing.Close(sp)

	cur, err := r.col.Aggregate(ctx, agg)

	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	var refreshToken model.RefreshTokenModel
	if cur.Next(ctx) {
		if err := cur.Decode(&refreshToken); err != nil {
			return nil, tracing.LogError(sp, err)
		}
	}

	tracing.LogResponse(sp, refreshToken)
	return refreshToken, nil
}
