package user

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryUser) InsertUser(ctx context.Context, in interface{}) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.InsertOne(ctx, in)
		})
	} else {
		res, err = r.col.InsertOne(ctx, in)
	}

	if err != nil {
		return tracing.LogError(sp, err)
	}

	tracing.LogResponse(sp, res)
	return nil
}
