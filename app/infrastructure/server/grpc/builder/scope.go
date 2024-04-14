package builder

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/scope"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

type ScopeBuilder struct{}

func NewScopeBuilder() scope.IScopeBuilder {
	return &ScopeBuilder{}
}

func (s *ScopeBuilder) GetScopeResponse(ctx context.Context, in interface{}) (*pb.GetScopeResponse, error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode data to proto format
	var data *model.ScopeModel
	if err := mapstructure.Decode(in, &data); err != nil {
		return nil, err
	}

	// convert endpoints object id to endpoints string
	//var endpoints []string
	//for _, endpoint := range data.Endpoints {
	//	endpoints = append(endpoints, endpoint.Hex())
	//}

	// manage endpoint
	var listEndpoint []*pb.EndpointData
	for _, endpoint := range data.ListEndpoint {
		listEndpoint = append(listEndpoint, &pb.EndpointData{
			Id:         endpoint.Id.Hex(),
			Service:    endpoint.Service,
			IsGenerate: *endpoint.IsGenerate,
			StillExist: *endpoint.StillExist,
			CreatedAt:  endpoint.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  endpoint.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// get custom code for response
	c := codes.Get(200)

	// manage response
	res := &pb.GetScopeResponse{
		Code:         int32(c.Code()),
		Message:      c.ResponseMessage,
		ErrorMessage: c.ErrorMessage,
		Data: &pb.ScopeData{
			Id:           data.Id.Hex(),
			Service:      data.Service,
			Name:         data.Name,
			ListEndpoint: listEndpoint,
			CreatedAt:    data.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    data.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (s *ScopeBuilder) GetScopesResponse(ctx context.Context, res, filter interface{}) (*pb.GetScopesResponse, error) {
	// tracing
	sp := tracing.StartChild(ctx, res, filter)
	defer tracing.Close(sp)

	// get custom code for response
	c := codes.Get(200)

	// manage data
	scopes, ok := res.([]model.ScopeModel)
	if !ok {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// manipulate data
	var scopesPb []*pb.ScopeData
	for _, val := range scopes {
		// convert endpoints object id to endpoints string
		var endpoints []string
		for _, endpoint := range val.Endpoints {
			endpoints = append(endpoints, endpoint.Hex())
		}

		// add to array
		scopesPb = append(scopesPb, &pb.ScopeData{
			Id:        val.Id.Hex(),
			Service:   val.Service,
			Name:      val.Name,
			Endpoints: endpoints,
			CreatedAt: val.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: val.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// get filter
	f, ok := filter.(model.ScopeFilter)
	if !ok {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// decode pagination
	var pagination *pb.Pagination
	if err := mapstructure.Decode(f.Pagination, &pagination); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// decode sort
	var sort []*pb.Sort
	if err := mapstructure.Decode(f.Sort.SortValue, &sort); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	return &pb.GetScopesResponse{
		Code:         int32(c.Code()),
		Message:      c.ResponseMessage,
		ErrorMessage: c.ErrorMessage,
		Data:         scopesPb,
		Pagination:   pagination,
		Sort:         sort,
	}, nil
}
