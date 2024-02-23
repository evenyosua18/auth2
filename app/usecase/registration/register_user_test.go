package registration

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
	"time"
)

func TestUsecaseRegistration_RegistrationUser(t *testing.T) {
	Convey("Test Registration User", t, func() {
		os.Setenv(constant.TokenDuration, "24")
		os.Setenv(constant.TokenSignature, "TestOauth2")

		// context
		ctx := context.Background()

		// repo
		userRepo := &authdb.MUserRepository{}
		accessTokenRepo := &authdb.MAccessTokenRepository{}
		refreshTokenRepo := &authdb.MRefreshTokenRepository{}

		Convey("Positive Scenario", func() {
			Convey("request valid and all repository works", func() {
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(nil, mongo.ErrNoDocuments).Once()
				userRepo.On("InsertUser", mock.Anything, mock.Anything).Return(nil).Once()

				// mock access token repository
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock refresh token repository
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldBeNil)
				})

				// result
				Convey("result checking", func() {
					So(res, ShouldNotBeNil)

					// decode response
					response := struct {
						Id           string
						RefreshToken string
						AccessToken  string
						ExpiredAt    int64
					}{}

					// decode
					err = mapstructure.Decode(res, &response)

					So(err, ShouldBeNil)
					So(response.ExpiredAt, ShouldNotBeZeroValue)
					So(response.AccessToken, ShouldNotBeEmpty)
					So(response.RefreshToken, ShouldNotBeEmpty)
				})
			})

			Convey("user already exist", func() {
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// set get user response
				getUserResponse := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  "TEST",
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(&getUserResponse, nil).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

				// error checking
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "410")
				})

				// response checking
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("invalid request", func() {
				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, "TEST")

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

			Convey("error when get user", func() {
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// error
				errRes := errors.New("TEST")

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(nil, errRes).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

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

			Convey("error when insert user", func() {
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// error
				errRes := errors.New("TEST")

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(nil, mongo.ErrNoDocuments).Once()
				userRepo.On("InsertUser", mock.Anything, mock.Anything).Return(errRes).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

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

			Convey("error when insert access token", func() {
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// error
				errRes := errors.New("TEST")

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(nil, mongo.ErrNoDocuments).Once()
				userRepo.On("InsertUser", mock.Anything, mock.Anything).Return(nil).Once()

				// mock access token repository
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

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
				// set request data
				req := struct {
					Username string
					Password string
					Email    string
					Phone    string
				}{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// error
				errRes := errors.New("TEST")

				// mock user repository
				userRepo.On("GetUser", mock.Anything, mock.Anything).Return(nil, mongo.ErrNoDocuments).Once()
				userRepo.On("InsertUser", mock.Anything, mock.Anything).Return(nil).Once()

				// mock access token repository
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock refresh token repository
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// new usecase
				uc := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

				// registration user
				res, err := uc.RegistrationUser(ctx, req)

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
