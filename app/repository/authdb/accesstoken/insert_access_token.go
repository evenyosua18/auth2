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

func (r *RepositoryAccessToken) InsertAccessToken(ctx context.Context, in interface{}) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// get model
	accessToken := model.AccessTokenModel{}.SetCreated()
	if err = mapstructure.Decode(in, &accessToken); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.InsertOne(ctx, accessToken)
		})
	} else {
		res, err = r.col.InsertOne(ctx, accessToken)
	}

	if err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return nil
}
