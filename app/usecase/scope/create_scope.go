package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type requestCreateScope struct {
	Name      string
	Service   string
	Endpoints []string
}

func (u *UsecaseScope) CreateScope(ctx context.Context, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode request
	var req requestCreateScope
	if err := mapstructure.Decode(in, &req); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	return u.scope.InsertScope(tracing.Context(sp), req.ToScopeModel(sp))
}

func (r requestCreateScope) ToScopeModel(sp interface{}) model.ScopeModel {
	// manage endpoint
	var endpoints []primitive.ObjectID

	for _, endpoint := range r.Endpoints {
		if endpointObjectId, err := primitive.ObjectIDFromHex(endpoint); err != nil {
			tracing.LogObject(sp, "endpoint_"+endpoint, err)
		} else {
			endpoints = append(endpoints, endpointObjectId)
		}
	}

	return model.ScopeModel{
		Service:   r.Service,
		Name:      r.Name,
		Endpoints: endpoints,
	}.SetCreated()
}
