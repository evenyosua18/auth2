package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryEndpoint) UpdateEndpoint(ctx context.Context, filter, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, filter, in)
	defer tracing.Close(sp)

	// decode filter
	var f model.EndpointFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// decode model
	var endpoint model.EndpointModel
	if err := mapstructure.Decode(in, &endpoint); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// update
	if res, err := r.col.UpdateOne(ctx, f.Filter(), endpoint.Update()); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 501))
	} else {
		tracing.LogResponse(sp, res)
		return nil
	}
}
