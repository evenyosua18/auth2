package accesstoken

import (
	"errors"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

func TestNewAccessTokenUsecase(t *testing.T) {
	Convey("Test Access Token Usecase Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock repository
			oauthClientRepo := &authdb.MOauthClientRepository{}
			userRepo := &authdb.MUserRepository{}
			accessTokenRepo := &authdb.MAccessTokenRepository{}
			refreshTokenRepo := &authdb.MRefreshTokenRepository{}

			res := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

			// response should be valid
			Convey("validate response should be valid", func() {
				So(res, ShouldEqual, &UsecaseAccessToken{
					oauthClient:  oauthClientRepo,
					user:         userRepo,
					accessToken:  accessTokenRepo,
					refreshToken: refreshTokenRepo,
				})
			})
		})
	})
}

func TestNewAccessTokenUsecase_ManageAccessToken(t *testing.T) {
	Convey("Test Manage Access Token", t, func() {
		os.Setenv(constant.TokenDuration, "24")
		os.Setenv(constant.TokenSignature, "TestOauth2")

		// mock repository
		oauthClientRepo := &authdb.MOauthClientRepository{}
		userRepo := &authdb.MUserRepository{}
		accessTokenRepo := &authdb.MAccessTokenRepository{}
		refreshTokenRepo := &authdb.MRefreshTokenRepository{}

		Convey("Positive Scenario", func() {
			Convey("request model valid and all repositories works", func() {
				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// user request
				user := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  "TEST",
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo).(*UsecaseAccessToken)

				// manage access token
				res, err := uc.manageAccessToken(nil, &user, 1)

				// error
				Convey("error should be nil", func() {
					So(err, ShouldBeNil)
				})

				// response
				Convey("response should not be nil", func() {
					So(res, ShouldNotBeNil)
				})

				// compare response
				Convey("compare with expected response", func() {
					// define res struct
					response := struct {
						AccessToken  string
						RefreshToken string
						ExpireAt     int64
					}{}

					// decode
					if err = mapstructure.Decode(res, &response); err != nil {
						t.Error(err)
					}

					So(response.RefreshToken, ShouldHaveLength, 16)
					So(response.AccessToken, ShouldNotBeEmpty)
					So(response.ExpireAt, ShouldBeGreaterThan, time.Now().Unix())
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("user model is empty", func() {
				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo).(*UsecaseAccessToken)

				// manage access token
				res, err := uc.manageAccessToken(nil, nil, 1)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "505")
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("insert access token is error", func() {
				// set error response
				errRes := errors.New("TEST")

				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// user request
				user := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  "TEST",
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo).(*UsecaseAccessToken)

				// manage access token
				res, err := uc.manageAccessToken(nil, &user, 1)

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

			Convey("insert refresh token is error", func() {
				// set error response
				errRes := errors.New("TEST")

				// mock insert access token
				accessTokenRepo.On("InsertAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(errRes).Once()

				// user request
				user := model.UserModel{
					Id:        primitive.NewObjectID(),
					Email:     "TEST",
					Phone:     "TEST",
					Username:  "TEST",
					Password:  "TEST",
					IsActive:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo).(*UsecaseAccessToken)

				// manage access token
				res, err := uc.manageAccessToken(nil, &user, 1)

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
