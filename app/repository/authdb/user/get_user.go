package user

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepositoryUser) GetUser(ctx context.Context, filter bson.M) (interface{}, error) {
	// start tracer
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// call database
	var res *model.UserModel
	if err := r.col.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
