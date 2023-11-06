package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryAccessToken) DeleteAccessToken(ctx context.Context, filter bson.M) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.DeleteOne(ctx, filter)
		})
	} else {
		res, err = r.col.DeleteOne(ctx, filter)
	}

	if err != nil {
		return tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return nil
}
