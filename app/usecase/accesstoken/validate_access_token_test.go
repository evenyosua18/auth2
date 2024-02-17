package accesstoken

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/utils/token"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestAccessTokenUsecase_ValidateAccessToken(t *testing.T) {
	// context
	ctx := context.Background()

	// init mock repository
	oauthClientRepo := &authdb.MOauthClientRepository{}
	userRepo := &authdb.MUserRepository{}
	accessTokenRepo := &authdb.MAccessTokenRepository{}
	refreshTokenRepo := &authdb.MRefreshTokenRepository{}

	Convey("Test Validate Access Token", t, func() {
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
					AccessToken string
				}{
					AccessToken: tok,
				}

				// set response get access token
				resAccessToken := model.AccessTokenModel{
					Id:        tokenId,
					ExpiredAt: expiredAt,
					UserId:    primitive.NewObjectID(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// mock get access token
				accessTokenRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return(&resAccessToken, nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// validate access token
				err = uc.ValidateAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Negative Scenario", func() {
			Convey("invalid request", func() {
				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// validate access token
				err := uc.ValidateAccessToken(ctx, "TEST")

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "502")
				})
			})

			Convey("error when get access token", func() {
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
					AccessToken string
				}{
					AccessToken: tok,
				}

				// set response get access token
				errRes := errors.New("TEST")

				// mock get access token
				accessTokenRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return(nil, errRes).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// validate access token
				err = uc.ValidateAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, errRes)
				})
			})

			Convey("token already expired", func() {
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
					AccessToken string
				}{
					AccessToken: tok,
				}

				// set response get access token
				resAccessToken := model.AccessTokenModel{
					Id:        tokenId,
					ExpiredAt: time.Now().Add(time.Hour * -1),
					UserId:    primitive.NewObjectID(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				// mock get access token
				accessTokenRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return(&resAccessToken, nil).Once()

				// create usecase
				uc := NewAccessTokenUsecase(oauthClientRepo, userRepo, accessTokenRepo, refreshTokenRepo)

				// validate access token
				err = uc.ValidateAccessToken(ctx, req)

				// error
				Convey("error checking", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldStartWith, "423")
				})
			})
		})
	})
}
