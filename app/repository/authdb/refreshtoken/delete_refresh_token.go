package refreshtoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (r *RepositoryRefreshToken) DeleteRefreshToken(ctx context.Context, filter bson.M) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.UpdateOne(ctx, filter, bson.M{
				"$set": bson.M{
					"deleted_at": time.Now(),
				},
			})
		})
	} else {
		res, err = r.col.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				"deleted_at": time.Now(),
			},
		})
	}

	if err != nil {
		return tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return nil
}
