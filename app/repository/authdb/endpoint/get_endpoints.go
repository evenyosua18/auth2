package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryEndpoint) GetEndpoints(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// decode filter
	var f model.EndpointFilter
	if err := mapstructure.Decode(filter, &f); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// db
	cur, err := r.col.Find(ctx, f.Filter())

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	// loop cursor
	var res []model.EndpointModel
	for cur.Next(ctx) {
		// decode row
		var endpoint model.EndpointModel
		if err := cur.Decode(&endpoint); err != nil {
			return nil, tracing.LogError(sp, codes.Wrap(err, 501))
		}

		// add to array
		res = append(res, endpoint)
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
