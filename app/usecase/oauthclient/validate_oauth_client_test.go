package oauthclient

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestOauthClientUsecase_ValidateOauthClient(t *testing.T) {
	Convey("Test Validate Oauth Client", t, func() {
		// context
		ctx := context.Background()

		//init repo
		oauthClientRepo := &authdb.MOauthClientRepository{}

		Convey("Positive Scenario", func() {
			Convey("valid request model and all repository works", func() {
				// client secret
				secret, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set request data
				req := struct {
					ClientId     string
					ClientSecret string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
				}

				// set response get oauth client
				getOauthClientResponse := &model.OauthClientModel{
					Id:           primitive.NewObjectID(),
					ClientId:     "TEST",
					ClientSecret: string(secret),
					Uri:          "TEST",
					Scopes:       "TEST",
					ClientType:   "TEST",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}

				// mock repository
				oauthClientRepo.On("GetOauthClient", mock.Anything, mock.Anything).Return(getOauthClientResponse, nil).Once()

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("invalid request model", func() {
				// request
				req := "TEST INVALID MODEL"

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should not be nil", func() {
					// error not nil
					So(err, ShouldNotBeNil)

					// error code
					So(err.Error(), ShouldStartWith, "502")
				})
			})

			Convey("error when get oauth client", func() {
				// set request data
				req := struct {
					ClientId     string
					ClientSecret string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
				}

				// set error
				errRes := errors.New("TEST")

				// mock repository
				oauthClientRepo.On("GetOauthClient", mock.Anything, mock.Anything).Return(nil, errRes).Once()

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should not be nil", func() {
					// error not nil
					So(err, ShouldNotBeNil)

					// compare error
					So(err, ShouldEqual, errRes)
				})
			})

			Convey("invalid response when get oauth client", func() {
				// set request data
				req := struct {
					ClientId     string
					ClientSecret string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
				}

				// mock repository
				oauthClientRepo.On("GetOauthClient", mock.Anything, mock.Anything).Return("TEST", nil).Once()

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should not be nil", func() {
					// error not nil
					So(err, ShouldNotBeNil)

					// error code
					So(err.Error(), ShouldStartWith, "502")
				})
			})

			Convey("requested client id is different", func() {
				// client secret
				secret, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set request data
				req := struct {
					ClientId     string
					ClientSecret string
				}{
					ClientId:     "TEST INVALID",
					ClientSecret: "TEST",
				}

				// set response get oauth client
				getOauthClientResponse := &model.OauthClientModel{
					Id:           primitive.NewObjectID(),
					ClientId:     "TEST",
					ClientSecret: string(secret),
					Uri:          "TEST",
					Scopes:       "TEST",
					ClientType:   "TEST",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}

				// mock repository
				oauthClientRepo.On("GetOauthClient", mock.Anything, mock.Anything).Return(getOauthClientResponse, nil).Once()

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should not be nil", func() {
					// error not nil
					So(err, ShouldNotBeNil)

					// error code
					So(err.Error(), ShouldStartWith, "421")
				})
			})

			Convey("invalid client secret", func() {
				// client secret
				secret, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set request data
				req := struct {
					ClientId     string
					ClientSecret string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST INVALID",
				}

				// set response get oauth client
				getOauthClientResponse := &model.OauthClientModel{
					Id:           primitive.NewObjectID(),
					ClientId:     "TEST",
					ClientSecret: string(secret),
					Uri:          "TEST",
					Scopes:       "TEST",
					ClientType:   "TEST",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}

				// mock repository
				oauthClientRepo.On("GetOauthClient", mock.Anything, mock.Anything).Return(getOauthClientResponse, nil).Once()

				// usecase
				uc := NewOauthClientUsecase(oauthClientRepo)

				// validate oauth client
				err := uc.ValidateOauthClient(ctx, req)

				// error
				Convey("error should not be nil", func() {
					// error not nil
					So(err, ShouldNotBeNil)

					// error code
					So(err.Error(), ShouldStartWith, "421")
				})
			})
		})
	})
}
