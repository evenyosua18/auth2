package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestScope(t *testing.T) {
	Convey("Test Scope Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				scope := ScopeModel{
					Service:   "TEST",
					Name:      "TEST",
					Endpoints: []primitive.ObjectID{primitive.NewObjectID()},
				}

				// set created
				newScope := scope.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should not be nil", func() {
						So(newScope.Id, ShouldNotEqual, primitive.NilObjectID)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newScope.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newScope.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newScope.DeletedAt, ShouldBeNil)
					})
				})
			})

			Convey("if id is not empty", func() {
				// set model
				scope := ScopeModel{
					Id:        primitive.NewObjectID(),
					Service:   "TEST",
					Name:      "TEST",
					Endpoints: []primitive.ObjectID{primitive.NewObjectID()},
				}

				// set created
				newScope := scope.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should be same", func() {
						So(newScope.Id, ShouldEqual, scope.Id)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newScope.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newScope.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newScope.DeletedAt, ShouldBeNil)
					})
				})
			})
		})

		// TODO Test Update
		Convey("Test Update", func() {
			Convey("update all existing value", func() {
				// set value
				id := primitive.NewObjectID()

				// set model
				scope := ScopeModel{
					Service:   "TEST",
					Name:      "TEST",
					Endpoints: []primitive.ObjectID{id},
					CreatedAt: time.Now().Add(time.Minute * -1),
					UpdatedAt: time.Now().Add(time.Minute * -1),
				}

				// update
				res := scope.Update()

				// result checking
				Convey("result checking", func() {
					// can't be empty
					Convey("result can't be empty", func() {
						So(res["$set"], ShouldNotBeEmpty)
					})

					// name
					Convey("value of name should be same", func() {
						So(res["$set"].(bson.M)["name"].(string), ShouldEqual, scope.Name)
					})

					// service
					Convey("value of service should be same", func() {
						So(res["$set"].(bson.M)["service"].(string), ShouldEqual, scope.Service)
					})

					// endpoints
					Convey("value of endpoints should be same", func() {
						So(res["$set"].(bson.M)["endpoints"].([]primitive.ObjectID), ShouldEqual, scope.Endpoints)
					})

					// updated at
					Convey("updated at should be greater than before", func() {
						So(res["$set"].(bson.M)["updated_at"].(time.Time).Unix(), ShouldBeGreaterThan, scope.UpdatedAt.Unix())
					})
				})

			})
		})

		Convey("Test Filter", func() {
			// set value
			id := primitive.NewObjectID()

			// set model
			filter := ScopeFilter{
				Id:        id.Hex(),
				Service:   "TEST",
				Name:      "TEST",
				Endpoints: []string{id.Hex()},
			}

			// expected result
			expectedFilter := bson.M{
				"deleted_at": nil,
				"_id":        id,
				"endpoints":  bson.M{"$in": []primitive.ObjectID{id}},
				"name":       primitive.Regex{Pattern: filter.Name, Options: "i"},
				"service":    filter.Service,
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
