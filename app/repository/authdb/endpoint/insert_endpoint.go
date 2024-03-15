package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryEndpoint) InsertEndpoint(ctx context.Context, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// get model
	endpoint := model.EndpointModel{}.SetCreated()
	if err := mapstructure.Decode(in, &endpoint); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	if res, err := r.col.InsertOne(ctx, endpoint); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, res)
		return nil
	}
}
