package accesstoken

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/utils/token"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

func TestAccessTokenUsecase_RefreshAccessToken(t *testing.T) {
	// set env
	os.Setenv(constant.TokenDuration, "24")
	os.Setenv(constant.TokenSignature, "TestOauth2")
	os.Setenv(constant.MaxRefreshToken, "3")

	// context
	ctx := context.Background()

	// init mock repository
	oauthClientRepo := &authdb.MOauthClientRepository{}
	userRepo := &authdb.MUserRepository{}
	accessTokenRepo := &authdb.MAccessTokenRepository{}
	refreshTokenRepo := &authdb.MRefreshTokenRepository{}

	Convey("Test Refresh Access Token", t, func() {
		Convey("Positive Scenario", func() {
			Convey("valid request and all repository works", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, expiredAt, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set refresh token response
				refreshTokenResp := model.RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: tokenId,
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(refreshTokenResp, nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock delete refresh token
				refreshTokenRepo.On("DeleteRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock update access token
				accessTokenRepo.On("UpdateAccessToken", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldBeNil)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldNotBeNil)

					// decode struct
					resStruct := struct {
						AccessToken  string
						RefreshToken string
						ExpireAt     int64
					}{}

					err := mapstructure.Decode(res, &resStruct)

					// decode should be works
					So(err, ShouldBeNil)
					So(resStruct.AccessToken, ShouldNotBeEmpty)
					So(resStruct.RefreshToken, ShouldNotEqual, refreshTokenResp.RefreshToken)
					So(resStruct.ExpireAt, ShouldEqual, expiredAt.Unix())
				})
			})

			Convey("return error if exceed maximum refresh times", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set refresh token response
				refreshTokenResp := model.RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: tokenId,
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         4,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(refreshTokenResp, nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "428")
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("invalid request", func() {
				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, "TEST")

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

			Convey("invalid refresh token response", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return("TEST", nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

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

			Convey("error when get refresh token", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set error response
				errResp := errors.New("TEST")

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(nil, errResp).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errResp)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("error when insert refresh token", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set refresh token response
				refreshTokenResp := model.RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: tokenId,
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				// set error response
				errResp := errors.New("TEST")

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(refreshTokenResp, nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(errResp).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errResp)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("error when delete refresh token", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set refresh token response
				refreshTokenResp := model.RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: tokenId,
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				// set error response
				errResp := errors.New("TEST")

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(refreshTokenResp, nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock delete refresh token
				refreshTokenRepo.On("DeleteRefreshToken", mock.Anything, mock.Anything).Return(errResp).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errResp)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})

			Convey("error when update access token", func() {
				// generate token id
				tokenId := primitive.NewObjectID()

				// create dummy token
				tok, _, err := token.GenerateToken(ctx, tokenId.Hex(), token.ClaimsInformation{
					Username: "TEST",
					Phone:    "TEST",
					Email:    "TEST",
				})

				if err != nil {
					t.Error(err)
				}

				// set request
				req := struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  tok,
					RefreshToken: "TEST",
				}

				// set refresh token response
				refreshTokenResp := model.RefreshTokenModel{
					Id:            primitive.NewObjectID(),
					AccessTokenId: tokenId,
					UserId:        primitive.NewObjectID(),
					RefreshToken:  "TEST",
					Count:         0,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				// set error response
				errResp := errors.New("TEST")

				// mock get refresh token
				refreshTokenRepo.On("GetRefreshToken", mock.Anything, mock.Anything).Return(refreshTokenResp, nil).Once()

				// mock insert refresh token
				refreshTokenRepo.On("InsertRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock delete refresh token
				refreshTokenRepo.On("DeleteRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()

				// mock update access token
				accessTokenRepo.On("UpdateAccessToken", mock.Anything, mock.Anything, mock.Anything).Return(errResp).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				res, err := uc.RefreshAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errResp)
				})

				// response
				Convey("response checking", func() {
					So(res, ShouldBeNil)
				})
			})
		})
	})
}
