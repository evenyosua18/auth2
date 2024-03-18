package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

type reportRegisterGRPC struct {
	TotalExistEndpoint      int
	ListNewEndpoints        []string
	ErrorCreateNewEndpoints []error
	ErrorUpdateEndpoints    []error
	ListIgnoreKeys          []string
	ListNotExistEndpoints   []string
}

func (u *UsecaseRegistrationEndpoint) RegisterGRPC(ctx context.Context, listEndpoint map[string]grpc.ServiceInfo) interface{} {
	// tracing
	sp := tracing.StartChild(ctx)
	defer tracing.Close(sp)

	// define reporting model
	report := reportRegisterGRPC{}

	// get current list endpoint
	listCurrentEndpointI, err := u.endpoint.GetEndpoints(ctx, model.EndpointFilter{
		Service: os.Getenv(constant.AppService),
	})

	if err != nil {
		log.Println(tracing.LogError(sp, err))
		return err
	}

	// type conversion
	listCurrentEndpoint, ok := listCurrentEndpointI.([]model.EndpointModel)

	if !ok {
		log.Println(tracing.LogError(sp, codes.Wrap(nil, 502)))
		return err
	}

	// map current endpoints
	mapCurrentEndpoints := make(map[string]*primitive.ObjectID, len(listCurrentEndpoint))

	for _, e := range listCurrentEndpoint {
		mapCurrentEndpoints[e.Endpoint] = &e.Id
	}

	// loop every next endpoint
	for key, value := range listEndpoint {
		// split key
		keys := strings.Split(key, ".")

		// ex: proto.AccessTokenService - split by ".", 0: validation, 1: save
		if len(keys) <= 1 || (len(keys) != 0 && u.IgnorePrefixes[keys[0]]) {
			report.ListIgnoreKeys = append(report.ListIgnoreKeys, key)
			continue
		}

		// loop all methods
		for _, method := range value.Methods {
			if mapCurrentEndpoints[keys[1]+"."+method.Name] != nil {
				// if key[1].method exists, delete from map current endpoints
				delete(mapCurrentEndpoints, keys[1]+"."+method.Name)
				report.TotalExistEndpoint++
			} else {
				// if key[1].method not exists, add new endpoint
				if err = u.endpoint.InsertEndpoint(ctx, model.EndpointModel{
					Service:    os.Getenv(constant.AppService),
					Endpoint:   keys[1] + "." + method.Name,
					IsGenerate: &constant.TrueValue,
					StillExist: &constant.TrueValue,
				}.SetCreated()); err != nil {
					report.ErrorCreateNewEndpoints = append(report.ErrorCreateNewEndpoints, err)
					continue
				} else {
					report.ListNewEndpoints = append(report.ListNewEndpoints, keys[1]+"."+method.Name)
				}
			}
		}
	}

	// loop for the rest list map endpoints, means the rest endpoints in this map is being deleted or not used anymore
	for endpointName, endpointId := range mapCurrentEndpoints {
		// update still exist value to false
		if err = u.endpoint.UpdateEndpoint(ctx, model.EndpointFilter{
			Id: endpointId,
		}, model.EndpointModel{
			StillExist: &constant.FalseValue,
		}); err != nil {
			report.ErrorUpdateEndpoints = append(report.ErrorUpdateEndpoints, err)
			continue
		} else {
			report.ListNotExistEndpoints = append(report.ListNotExistEndpoints, endpointName)
		}
	}

	tracing.LogResponse(sp, report)
	return report
}
