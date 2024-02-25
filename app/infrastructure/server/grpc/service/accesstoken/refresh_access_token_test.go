package accesstoken

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/mock/builder"
	"github.com/evenyosua18/auth2/app/mock/usecase"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceAccessToken_RefreshAccessToken(t *testing.T) {
	Convey("Test Refresh Access Token", t, func() {
		// ctx
		ctx := context.Background()

		// mock
		out := &builder.MAccessTokenBuilder{}
		uc := &usecase.MAccessTokenUsecase{}

		Convey("Positive Scenario", func() {
			// set request
			req := &pb.RefreshTokenRequest{
				AccessToken:  "TEST",
				RefreshToken: "TEST",
			}

			// usecase response
			resUC := struct {
				AccessToken  string
				RefreshToken string
				ExpiredAt    int64
			}{
				AccessToken:  "TEST",
				RefreshToken: "TEST",
				ExpiredAt:    0,
			}

			// builder response
			resBuilder := &pb.AccessTokenResponse{
				Code:         0,
				Message:      "TEST",
				ErrorMessage: "TESt",
				Data:         nil,
			}

			// mock usecase
			uc.On("RefreshAccessToken", mock.Anything, mock.Anything).Return(&resUC, nil).Once()

			// mock builder
			out.On("AccessTokenResponse", mock.Anything, mock.Anything, mock.Anything).Return(resBuilder, nil).Once()

			// create service
			svc := NewAccessTokenService(out, uc)

			// refresh access token
			res, err := svc.RefreshAccessToken(ctx, req)

			// error
			Convey("error checking", func() {
				So(err, ShouldBeNil)
			})

			// result
			Convey("result checking", func() {
				So(res, ShouldNotBeNil)
			})
		})

		Convey("Negative Scenario", func() {
			// set request
			req := &pb.RefreshTokenRequest{
				AccessToken:  "TEST",
				RefreshToken: "TEST",
			}

			// error
			errRes := errors.New("TEST")

			// mock usecase
			uc.On("RefreshAccessToken", mock.Anything, mock.Anything).Return(nil, errRes).Once()

			// create service
			svc := NewAccessTokenService(out, uc)

			// refresh access token
			res, err := svc.RefreshAccessToken(ctx, req)

			// error
			Convey("error checking", func() {
				So(err, ShouldNotBeNil)
			})

			// result
			Convey("result checking", func() {
				So(res, ShouldBeNil)
			})
		})
	})
}
