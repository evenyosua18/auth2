package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestEndpoint(t *testing.T) {
	Convey("Test Endpoint Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				endpoint := EndpointModel{
					Service:  "TEST",
					Endpoint: "TEST",
				}

				// set created
				newEndpoint := endpoint.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should not be nil", func() {
						So(newEndpoint.Id, ShouldNotEqual, primitive.NilObjectID)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newEndpoint.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newEndpoint.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newEndpoint.DeletedAt, ShouldBeNil)
					})
				})
			})

			Convey("if id is not empty", func() {
				// set model
				endpoint := EndpointModel{
					Id:       primitive.NewObjectID(),
					Service:  "TEST",
					Endpoint: "TEST",
				}

				// set created
				newEndpoint := endpoint.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should be same", func() {
						So(newEndpoint.Id, ShouldEqual, endpoint.Id)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newEndpoint.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newEndpoint.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newEndpoint.DeletedAt, ShouldBeNil)
					})
				})
			})
		})

		Convey("Test Update", func() {
			Convey("Update all existing value", func() {
				// const value
				trueValue := true

				// set model
				endpoint := EndpointModel{
					Service:    "TEST",
					Endpoint:   "TEST",
					IsGenerate: &trueValue,
					StillExist: &trueValue,
					CreatedAt:  time.Now().Add(time.Minute * -1),
					UpdatedAt:  time.Now().Add(time.Minute * -1),
				}

				// update
				res := endpoint.Update()

				// result checking
				Convey("result checking", func() {
					// can't be empty
					Convey("result can't be empty", func() {
						So(res["$set"], ShouldNotBeEmpty)
					})

					// service
					Convey("service should be same", func() {
						So(res["$set"].(bson.M)["service"], ShouldEqual, endpoint.Service)
					})

					// endpoint
					Convey("endpoint should be same", func() {
						So(res["$set"].(bson.M)["endpoint"], ShouldEqual, endpoint.Endpoint)
					})

					// is generated
					Convey("is generate value should be same", func() {
						So(res["$set"].(bson.M)["is_generate"], ShouldEqual, endpoint.IsGenerate)
					})

					// still exist
					Convey("still exist value should be same", func() {
						So(res["$set"].(bson.M)["still_exist"], ShouldEqual, endpoint.StillExist)
					})

					// updated at
					Convey("updated at should be greater than before", func() {
						So(res["$set"].(bson.M)["updated_at"].(time.Time).Unix(), ShouldBeGreaterThan, endpoint.UpdatedAt.Unix())
					})
				})
			})
		})

		Convey("Test Filter", func() {
			// const value
			trueValue := true
			id := primitive.NewObjectID()

			// set model
			filter := EndpointFilter{
				Id:         &id,
				Service:    "TEST",
				Endpoint:   "TEST",
				IsGenerate: &trueValue,
				StillExist: &trueValue,
			}

			// expected result
			expectedFilter := bson.M{
				"deleted_at":  nil,
				"_id":         &id,
				"endpoint":    primitive.Regex{Pattern: "TEST", Options: "i"},
				"is_generate": &trueValue,
				"service":     "TEST",
				"still_exist": &trueValue,
			}

			// filter
			res := filter.Filter()

			// result
			Convey("filter should be same with the expected", func() {
				So(res, ShouldEqual, expectedFilter)
			})
		})
	})
}
