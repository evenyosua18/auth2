package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestOauthClient(t *testing.T) {
	Convey("Test Oauth Client Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				oauthClient := OauthClientModel{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Uri:          "TEST",
					Scopes:       "TEST",
					ClientType:   "TEST",
				}

				// created
				res := oauthClient.SetCreated()

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
				oauthClient := OauthClientModel{
					Id:           primitive.NewObjectID(),
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Uri:          "TEST",
					Scopes:       "TEST",
					ClientType:   "TEST",
				}

				// created
				res := oauthClient.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should not be empty", func() {
						So(res.Id, ShouldEqual, oauthClient.Id)
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
			// set value
			id := primitive.NewObjectID()

			// set model
			filter := OauthClientFilter{
				Id:       &id,
				ClientId: "TEST",
			}

			// set expected
			expectedFilter := bson.M{
				"deleted_at": nil,
				"_id":        &id,
				"client_id":  filter.ClientId,
			}

			// get filter
			res := filter.Filter()

			// result
			Convey("filter should be same with the expected", func() {
				So(res, ShouldEqual, expectedFilter)
			})
		})
	})
}
