package accesstoken

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

func TestAccessTokenUsecase_PasswordGrant(t *testing.T) {
	// context
	ctx := context.Background()

	// init mock repository
	oauthClientRepo := &authdb.MOauthClientRepository{}
	userRepo := &authdb.MUserRepository{}
	accessTokenRepo := &authdb.MAccessTokenRepository{}
	refreshTokenRepo := &authdb.MRefreshTokenRepository{}

	Convey("Test Password Grant", t, func() {
		Convey("Positive Scenario", func() {
			Convey("invalid password", func() {
				// set request
				req := struct {
					ClientId     string
					ClientSecret string
					Username     string
					Password     string
					Scopes       string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Username:     "TEST",
					Password:     "TEST INVALID",
					Scopes:       "TEST",
				}

				// set user password
				pw, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set response get user
				resUser := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  string(pw),
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// mock get user
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(&resUser, nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "402")
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("valid password and all repository works", func() {
				// set request
				req := struct {
					ClientId     string
					ClientSecret string
					Username     string
					Password     string
					Scopes       string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Username:     "TEST",
					Password:     "TEST",
					Scopes:       "TEST",
				}

				// set user password
				pw, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set response get user
				resUser := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  string(pw),
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// mock get user
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(&resUser, nil).Once()

				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldBeNil)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldNotBeNil)
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("invalid request", func() {
				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, "TEST")

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "502")
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("invalid response from get user", func() {
				// set request
				req := struct {
					ClientId     string
					ClientSecret string
					Username     string
					Password     string
					Scopes       string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Username:     "TEST",
					Password:     "TEST",
					Scopes:       "TEST",
				}

				// mock get user
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return("TEST INVALID REQUEST", nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "502")
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("error when insert access token", func() {
				// set request
				req := struct {
					ClientId     string
					ClientSecret string
					Username     string
					Password     string
					Scopes       string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Username:     "TEST",
					Password:     "TEST",
					Scopes:       "TEST",
				}

				// set user password
				pw, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set response get user
				resUser := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  string(pw),
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// set error response
				errRes := errors.New("TEST")

				// mock get user
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(&resUser, nil).Once()

				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errRes)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("error when insert refresh token", func() {
				// set request
				req := struct {
					ClientId     string
					ClientSecret string
					Username     string
					Password     string
					Scopes       string
				}{
					ClientId:     "TEST",
					ClientSecret: "TEST",
					Username:     "TEST",
					Password:     "TEST",
					Scopes:       "TEST",
				}

				// set user password
				pw, _ := bcrypt.GenerateFromPassword([]byte("TEST"), bcrypt.DefaultCost)

				// set response get user
				resUser := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  string(pw),
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// set error response
				errRes := errors.New("TEST")

				// mock get user
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(&resUser, nil).Once()

				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// password grant
				res, err := uc.PasswordGrant(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errRes)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})
		})
	})
}
