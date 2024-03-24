package registration

import (
	"context"
	"errors"
	"fmt"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/codes"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestUsecaseRegistrationEndpoint_RegisterGRPC(t *testing.T) {
	// NOTE: ignored service means service name starts with value that listing on map ignored service
	Convey("Test Registration GRPC", t, func() {
		// context
		ctx := context.Background()

		// repo
		endpointRepo := &authdb.MEndpointRepository{}

		Convey("Positive Scenario", func() {
			// mock insert endpoint
			endpointRepo.On("InsertEndpoint", mock.Anything, mock.Anything).Return(nil).Maybe()

			// mock update endpoint
			endpointRepo.On("UpdateEndpoint", mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

			// five existing endpoints
			// 10 request endpoints (3 service with 3 exist, 4 new endpoint and 2 not exist anymore, 2 ignored service)
			Convey("specific case with existing endpoints", func() {
				// true value
				trueValue := true

				// expected value
				const (
					totalExist    = 3
					totalNew      = 4
					totalNotExist = 2
					totalIgnored  = 2
				)

				// list endpoints, service1: 1,2,3, service2: 1,2
				listExistingEndpoints := []model.EndpointModel{
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-1",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-2",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-3",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service2.test-endpoint-1",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service2.test-endpoint-2",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
				}

				// mock get list endpoints
				endpointRepo.On("GetEndpoints", mock.Anything, mock.Anything).Return(listExistingEndpoints, nil).Once()

				// list ongoing endpoints
				listOngoingEndpoints := map[string]grpc.ServiceInfo{
					"proto.service1": {
						Methods: []grpc.MethodInfo{
							{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-2", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-3", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-4", IsClientStream: true, IsServerStream: true}, // new
							{Name: "test-endpoint-5", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"proto.service2": {
						Methods: []grpc.MethodInfo{
							//{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // not exist anymore
							//{Name: "test-endpoint-2", IsClientStream: true, IsServerStream: true}, // not exist anymore
							{Name: "test-endpoint-3", IsClientStream: true, IsServerStream: true}, // new
							{Name: "test-endpoint-4", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"grpc.service1": {},
					"test":          {},
				}

				// create usecase
				uc := NewRegistrationEndpointUsecase(endpointRepo)

				// register grpc
				res := uc.RegisterGRPC(ctx, listOngoingEndpoints)

				// check result
				Convey("result checking", func() {
					// get res
					resObj := res.(reportRegisterGRPC)

					// total exist endpoint
					Convey(fmt.Sprintf("total exist endpoints should be %d", totalExist), func() {
						So(resObj.TotalExistEndpoint, ShouldEqual, totalExist)
					})

					// total new endpoint
					Convey(fmt.Sprintf("total new endpoints should be %d", totalNew), func() {
						So(len(resObj.ListNewEndpoints), ShouldEqual, totalNew)
					})

					// total ignored keys
					Convey(fmt.Sprintf("total ignored service should be %d", totalIgnored), func() {
						So(len(resObj.ListIgnoreKeys), ShouldEqual, totalIgnored)
					})

					// total not exist endpoints
					Convey(fmt.Sprintf("total not exist endpoints should be %d", totalNotExist), func() {
						So(len(resObj.ListNotExistEndpoints), ShouldEqual, totalNotExist)
					})

					// total error
					Convey("total error should be zero", func() {
						So(len(resObj.ErrorCreateNewEndpoints)+len(resObj.ErrorUpdateEndpoints), ShouldBeZeroValue)
					})
				})
			})

			// there is no existing endpoint
			// 3 request endpoints (2 service with 3 new endpoint and 2 ignored service)
			Convey("specific case without existing endpoints", func() {
				// expected value
				const (
					totalExist    = 0
					totalNew      = 2
					totalNotExist = 0
					totalIgnored  = 2
				)

				// list endpoints, empty
				var listExistingEndpoints []model.EndpointModel

				// mock insert endpoint
				endpointRepo.On("InsertEndpoint", mock.Anything, mock.Anything).Return(nil).Maybe()

				// mock update endpoint
				endpointRepo.On("UpdateEndpoint", mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

				// list ongoing endpoints
				listOngoingEndpoints := map[string]grpc.ServiceInfo{
					"proto.service1": {
						Methods: []grpc.MethodInfo{
							{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"proto.service2": {
						Methods: []grpc.MethodInfo{
							{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"grpc.service1": {},
					"test":          {},
				}

				// mock get list endpoints
				endpointRepo.On("GetEndpoints", mock.Anything, mock.Anything).Return(listExistingEndpoints, nil).Once()

				// create usecase
				uc := NewRegistrationEndpointUsecase(endpointRepo)

				// register grpc
				res := uc.RegisterGRPC(ctx, listOngoingEndpoints)

				// check result
				Convey("result checking", func() {
					// get res
					resObj := res.(reportRegisterGRPC)

					// total exist endpoint
					Convey(fmt.Sprintf("total exist endpoints should be %d", totalExist), func() {
						So(resObj.TotalExistEndpoint, ShouldEqual, totalExist)
					})

					// total new endpoint
					Convey(fmt.Sprintf("total new endpoints should be %d", totalNew), func() {
						So(len(resObj.ListNewEndpoints), ShouldEqual, totalNew)
					})

					// total ignored keys
					Convey(fmt.Sprintf("total ignored service should be %d", totalIgnored), func() {
						So(len(resObj.ListIgnoreKeys), ShouldEqual, totalIgnored)
					})

					// total not exist endpoints
					Convey(fmt.Sprintf("total not exist endpoints should be %d", totalNotExist), func() {
						So(len(resObj.ListNotExistEndpoints), ShouldEqual, totalNotExist)
					})

					// total error
					Convey("total error should be zero", func() {
						So(len(resObj.ErrorCreateNewEndpoints)+len(resObj.ErrorUpdateEndpoints), ShouldBeZeroValue)
					})
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("error when get endpoints", func() {
				// error response
				errorRes := errors.New("TEST")

				// list ongoing endpoints
				listOngoingEndpoints := map[string]grpc.ServiceInfo{}

				// mock get list endpoints
				endpointRepo.On("GetEndpoints", mock.Anything, mock.Anything).Return(nil, errorRes).Once()

				// usecase
				uc := NewRegistrationEndpointUsecase(endpointRepo)

				// register grpc
				res := uc.RegisterGRPC(ctx, listOngoingEndpoints)

				// result checking
				Convey("result checking", func() {
					Convey("response should be error", func() {
						So(res, ShouldEqual, errorRes)
					})
				})
			})

			Convey("invalid list endpoints", func() {
				// list ongoing endpoints
				listOngoingEndpoints := map[string]grpc.ServiceInfo{}

				// mock get list endpoints
				endpointRepo.On("GetEndpoints", mock.Anything, mock.Anything).Return("invalid model", nil).Once()

				// usecase
				uc := NewRegistrationEndpointUsecase(endpointRepo)

				// register grpc
				res := uc.RegisterGRPC(ctx, listOngoingEndpoints)

				// result checking
				Convey("result checking", func() {
					Convey("response should be error", func() {
						So(res, ShouldEqual, codes.Wrap(nil, 502))
					})
				})
			})

			// five existing endpoints
			// 10 request endpoints (3 service with 3 exist, 4 new endpoint and 2 not exist anymore, 2 ignored service)
			Convey("error when update and insert new endpoint", func() {
				// true value
				trueValue := true

				// expected value
				const (
					totalExist       = 3
					totalNew         = 0
					totalNotExist    = 0
					totalCreateError = 4
					totalUpdateError = 2
					totalIgnored     = 2
				)

				// list endpoints, service1: 1,2,3, service2: 1,2
				listExistingEndpoints := []model.EndpointModel{
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-1",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-2",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service1.test-endpoint-3",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service2.test-endpoint-1",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
					{
						Id:         primitive.NewObjectID(),
						Service:    "auth2",
						Endpoint:   "service2.test-endpoint-2",
						IsGenerate: &trueValue,
						StillExist: &trueValue,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					},
				}

				// mock get list endpoints
				endpointRepo.On("GetEndpoints", mock.Anything, mock.Anything).Return(listExistingEndpoints, nil).Once()

				// mock insert endpoint
				endpointRepo.On("InsertEndpoint", mock.Anything, mock.Anything).Return(errors.New("TEST")).Maybe()

				// mock update endpoint
				endpointRepo.On("UpdateEndpoint", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("TEST")).Maybe()

				// list ongoing endpoints
				listOngoingEndpoints := map[string]grpc.ServiceInfo{
					"proto.service1": {
						Methods: []grpc.MethodInfo{
							{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-2", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-3", IsClientStream: true, IsServerStream: true}, // exist
							{Name: "test-endpoint-4", IsClientStream: true, IsServerStream: true}, // new
							{Name: "test-endpoint-5", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"proto.service2": {
						Methods: []grpc.MethodInfo{
							//{Name: "test-endpoint-1", IsClientStream: true, IsServerStream: true}, // not exist anymore
							//{Name: "test-endpoint-2", IsClientStream: true, IsServerStream: true}, // not exist anymore
							{Name: "test-endpoint-3", IsClientStream: true, IsServerStream: true}, // new
							{Name: "test-endpoint-4", IsClientStream: true, IsServerStream: true}, // new
						},
					},
					"grpc.service1": {},
					"test":          {},
				}

				// create usecase
				uc := NewRegistrationEndpointUsecase(endpointRepo)

				// register grpc
				res := uc.RegisterGRPC(ctx, listOngoingEndpoints)

				// check result
				Convey("result checking", func() {
					// get res
					resObj := res.(reportRegisterGRPC)

					// total exist endpoint
					Convey(fmt.Sprintf("total exist endpoints should be %d", totalExist), func() {
						So(resObj.TotalExistEndpoint, ShouldEqual, totalExist)
					})

					// total new endpoint
					Convey(fmt.Sprintf("total new endpoints should be %d", totalNew), func() {
						So(len(resObj.ListNewEndpoints), ShouldEqual, totalNew)
					})

					// total ignored keys
					Convey(fmt.Sprintf("total ignored service should be %d", totalIgnored), func() {
						So(len(resObj.ListIgnoreKeys), ShouldEqual, totalIgnored)
					})

					// total not exist endpoints
					Convey(fmt.Sprintf("total not exist endpoints should be %d", totalNotExist), func() {
						So(len(resObj.ListNotExistEndpoints), ShouldEqual, totalNotExist)
					})

					// total create error
					Convey(fmt.Sprintf("total error create endpoints should be %d", totalCreateError), func() {
						So(len(resObj.ErrorCreateNewEndpoints), ShouldEqual, totalCreateError)
					})

					// total update error
					Convey(fmt.Sprintf("total error create endpoints should be %d", totalUpdateError), func() {
						So(len(resObj.ErrorUpdateEndpoints), ShouldEqual, totalUpdateError)
					})
				})
			})
		})
	})
}
