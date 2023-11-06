package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepositoryAccessToken) GetAccessToken(ctx context.Context, filter bson.M) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// call database
	var res *model.AccessTokenModel
	if err := r.col.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
