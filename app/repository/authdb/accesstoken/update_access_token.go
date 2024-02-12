package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryAccessToken) UpdateAccessToken(ctx context.Context, filter, in interface{}) (res interface{}, err error) {
	// tracing
	sp := tracing.StartChild(ctx, in, filter)
	defer tracing.Close(sp)

	// filter
	var f model.AccessTokenFilter
	if err = mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// model
	var accessToken model.AccessTokenModel
	if err = mapstructure.Decode(in, &accessToken); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.UpdateOne(ctx, f.Filter(), accessToken.Update())
		})
	} else {
		res, err = r.col.UpdateOne(ctx, f.Filter(), accessToken.Update())
	}

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
