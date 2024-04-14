package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UsecaseScope) GetScope(ctx context.Context, filter interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	// get scope
	res, err := u.scope.GetScope(tracing.Context(sp), filter)
	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	// convert to scope model
	scope, ok := res.(model.ScopeModel)

	if !ok {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// get list endpoints
	if endpoints, err := u.endpoint.GetEndpoints(tracing.Context(sp), struct {
		ListId []primitive.ObjectID
	}{
		ListId: scope.Endpoints,
	}); err != nil {
		return nil, tracing.LogError(sp, err)
	} else {
		// convert to endpoints model
		scope.ListEndpoint = endpoints.([]model.EndpointModel)
	}

	tracing.LogResponse(sp, scope)
	return scope, nil
}
