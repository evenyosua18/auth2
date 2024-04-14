package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
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

		// TODO Test Filter
		//Convey("Test Filter")

		// TODO Test Update
		//Convey("Test Update")
	})
}
