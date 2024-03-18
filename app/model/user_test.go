package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUser(t *testing.T) {
	Convey("Test User Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				user := UserModel{
					Email:    "TEST",
					Phone:    "TEST",
					Username: "TEST",
					Password: "TEST",
					IsActive: true,
				}

				// set created
				res := user.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should not be empty", func() {
						So(res.Id, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(res.CreatedAt, ShouldNotBeEmpty)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(res.UpdatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(res.DeletedAt, ShouldBeNil)
					})
				})
			})

			Convey("if id is not empty", func() {
				// set model
				user := UserModel{
					Id:       primitive.NewObjectID(),
					Email:    "TEST",
					Phone:    "TEST",
					Username: "TEST",
					Password: "TEST",
					IsActive: true,
				}

				// set created
				res := user.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should be same", func() {
						So(res.Id, ShouldEqual, user.Id)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(res.CreatedAt, ShouldNotBeEmpty)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(res.UpdatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(res.DeletedAt, ShouldBeNil)
					})
				})
			})
		})

		Convey("Test Filter", func() {
			Convey("check exist", func() {
				// set value
				id := primitive.NewObjectID()

				// set model
				filter := UserFilter{
					Id:       &id,
					Email:    "TEST",
					Username: "TEST",
					Phone:    "TEST",
					IsActive: true,
					Value:    "TEST",
				}

				// set expected
				expectedFilter := bson.M{
					"deleted_at": nil,
					"_id":        &id,
					"is_active":  true,
					"phone":      primitive.Regex{Pattern: filter.Phone, Options: "i"},
					"email":      primitive.Regex{Pattern: filter.Email, Options: "i"},
					"username":   filter.Username,
					"$or":        bson.A{bson.M{"username": filter.Username}, bson.M{"email": filter.Email}, bson.M{"phone": filter.Phone}},
				}

				// filter
				res := filter.Filter()

				// result
				Convey("filter should be same with the expected", func() {
					So(res, ShouldEqual, expectedFilter)
				})
			})

			Convey("not check exist", func() {
				// set value
				id := primitive.NewObjectID()

				// set model
				filter := UserFilter{
					Id:         &id,
					Email:      "TEST",
					Username:   "TEST",
					Phone:      "TEST",
					IsActive:   true,
					CheckExist: true,
				}

				// set expected
				expectedFilter := bson.M{
					"deleted_at": nil,
					"_id":        &id,
					"is_active":  true,
					"$or":        bson.A{bson.M{"phone": filter.Phone}, bson.M{"email": filter.Email}, bson.M{"username": filter.Username}},
				}

				// filter
				res := filter.Filter()

				// result
				Convey("filter should be same with the expected", func() {
					So(res, ShouldEqual, expectedFilter)
				})
			})
		})
	})
}
