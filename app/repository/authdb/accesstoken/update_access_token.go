package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryAccessToken) UpdateAccessToken(ctx context.Context, filter, in bson.M) (res interface{}, err error) {
	// tracing
	sp := tracing.StartChild(ctx, in, filter)
	defer tracing.Close(sp)

	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.UpdateOne(ctx, filter, bson.M{
				"$set": in,
			})
		})
	} else {
		res, err = r.col.UpdateOne(ctx, filter, bson.M{
			"$set": in,
		})
	}

	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
