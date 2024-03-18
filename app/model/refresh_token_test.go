package model

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	// set current default array
	defaultQuery := 3

	Convey("Test Refresh Token Model", t, func() {
		Convey("Test SetCreated", func() {
			Convey("if id is empty", func() {
				// set model
				refreshToken := RefreshTokenModel{
					AccessTokenId: primitive.NewObjectID(),
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
				}

				// set created
				res := refreshToken.SetCreated()

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
				refreshToken := RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: primitive.NewObjectID(),
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
				}

				// set created
				res := refreshToken.SetCreated()

				// result
				Convey("result checking", func() {
					// id
					Convey("id should be same", func() {
						So(res.Id, ShouldEqual, refreshToken.Id)
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
			filter := RefreshTokenFilter{
				Id:           &id,
				RefreshToken: "TEST",
			}

			// set expected
			expectedFilter := bson.M{
				"deleted_at":    nil,
				"_id":           &id,
				"refresh_token": filter.RefreshToken,
			}

			// filter
			res := filter.Filter()

			// result
			Convey("filter should be same with the expected", func() {
				So(res, ShouldEqual, expectedFilter)
			})
		})

		Convey("Test Aggregator", func() {
			Convey("without addition", func() {
				// set value
				id := primitive.NewObjectID()

				// set filter
				filter := RefreshTokenFilter{
					Id:           &id,
					RefreshToken: "TEST",
				}

				// get filter result
				resFilter := bson.M{"$match": filter.Filter()}

				// set aggregator
				agg := filter.Aggregate()

				// result
				Convey("result checking", func() {
					Convey("should not be empty", func() {
						So(agg, ShouldNotBeEmpty)
					})

					Convey(fmt.Sprintf("length of aggregate should be %d", defaultQuery), func() {
						So(len(agg), ShouldEqual, defaultQuery)
					})

					if len(agg) > 0 {
						Convey("the first value should be same with filter", func() {
							So(agg[0], ShouldEqual, resFilter)
						})
					}
				})
			})

			Convey("with additions", func() {
				// set value
				id := primitive.NewObjectID()

				// set filter
				filter := RefreshTokenFilter{
					Id:           &id,
					RefreshToken: "TEST",
				}

				// get filter result
				resFilter := bson.M{"$match": filter.Filter()}

				// set additions
				adds := []bson.M{
					{
						"first": "test",
					},
					{
						"second": "test",
					},
				}

				// set aggregator
				agg := filter.Aggregate(adds...)

				// result
				Convey("result checking", func() {
					Convey("should not be empty", func() {
						So(agg, ShouldNotBeEmpty)
					})

					Convey(fmt.Sprintf("length of aggregate should be %d", defaultQuery+len(adds)), func() {
						So(len(agg), ShouldEqual, defaultQuery+len(adds))
					})

					if len(agg) == defaultQuery+len(adds) {
						Convey("the first value should be same with filter", func() {
							So(agg[0], ShouldEqual, resFilter)
						})

						Convey("the second array till length of additions + 1 should be same", func() {
							So(agg[1:len(adds)+1], ShouldEqual, adds)
						})
					}
				})
			})
		})
	})
}
