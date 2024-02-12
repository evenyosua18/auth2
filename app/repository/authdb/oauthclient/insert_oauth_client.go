package oauthclient

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *RepositoryOauthClient) InsertOauthClient(ctx context.Context, in interface{}) (err error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// model
	oauthClient := model.OauthClientModel{}.SetCreated()
	if err = mapstructure.Decode(in, &oauthClient); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	var res interface{}
	if ss := r.db.GetSession(ctx.Value(constant.SessionId)); ss != nil {
		tracing.LogObject(sp, "use session", true)
		res, err = ss.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
			return r.col.InsertOne(ctx, oauthClient)
		})
	} else {
		res, err = r.col.InsertOne(ctx, oauthClient)
	}

	if err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	}

	tracing.LogResponse(sp, res)
	return nil
}
