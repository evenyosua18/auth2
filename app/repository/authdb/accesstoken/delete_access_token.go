package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryAccessToken) DeleteAccessToken(ctx context.Context, filter interface{}) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// get filter
	var f model.AccessTokenFilter
	if err = mapstructure.Decode(filter, &f); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.DeleteOne(ctx, f.Filter())
		})
	} else {
		res, err = r.col.DeleteOne(ctx, f.Filter())
	}

	if err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return nil
}
