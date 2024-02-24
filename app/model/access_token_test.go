package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestAccessToken(t *testing.T) {
	Convey("Test Access Token Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				accessToken := AccessTokenModel{
					ExpiredAt: time.Now().Add(time.Hour * 24),
					UserId:    primitive.NewObjectID(),
				}

				// set created
				newAccessToken := accessToken.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should not be nil", func() {
						So(newAccessToken.Id, ShouldNotEqual, primitive.NilObjectID)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newAccessToken.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newAccessToken.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newAccessToken.DeletedAt, ShouldBeNil)
					})
				})
			})

			Convey("if id is not empty", func() {
				// set model
				accessToken := AccessTokenModel{
					Id:        primitive.NewObjectID(),
					ExpiredAt: time.Now().Add(time.Hour * 24),
					UserId:    primitive.NewObjectID(),
				}

				// set created
				newAccessToken := accessToken.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should be same", func() {
						So(newAccessToken.Id, ShouldEqual, accessToken.Id)
					})

					// updated at
					Convey("updated at should not be empty", func() {
						So(newAccessToken.UpdatedAt, ShouldNotBeEmpty)
					})

					// created at
					Convey("created at should not be empty", func() {
						So(newAccessToken.CreatedAt, ShouldNotBeEmpty)
					})

					// deleted at
					Convey("deleted at should be nil", func() {
						So(newAccessToken.DeletedAt, ShouldBeNil)
					})
				})
			})

		})

		Convey("Test Update", func() {
			Convey("update all existing value", func() {
				// set model
				accessToken := AccessTokenModel{
					Id:        primitive.NewObjectID(),
					ExpiredAt: time.Now(),
					UserId:    primitive.NewObjectID(),
					CreatedAt: time.Now().Add(time.Minute * -1),
					UpdatedAt: time.Now().Add(time.Minute * -1),
				}

				// update
				res := accessToken.Update()

				// result checking
				Convey("result checking", func() {
					// can't be empty
					Convey("result can't be empty", func() {
						So(res["$set"], ShouldNotBeEmpty)
					})

					// expired at
					Convey("expired at should be same", func() {
						So(res["$set"].(bson.M)["expired_at"], ShouldEqual, accessToken.ExpiredAt)
					})

					// updated at
					Convey("updated at should be greater than before", func() {
						So(res["$set"].(bson.M)["updated_at"].(time.Time).Unix(), ShouldBeGreaterThan, accessToken.UpdatedAt.Unix())
					})
				})
			})
		})

		Convey("Test Filter", func() {
			// set value
			id := primitive.NewObjectID()
			userId := primitive.NewObjectID()
			expiredAt := time.Now()

			// set model
			filter := AccessTokenFilter{
				Id:        &id,
				UserId:    &userId,
				ExpiredAt: &expiredAt,
			}

			// expected
			expectedFilter := bson.M{
				"deleted_at": nil,
				"_id":        &id,
				"user_id":    &userId,
				"expired_at": bson.M{
					"$gte": &expiredAt,
				},
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
