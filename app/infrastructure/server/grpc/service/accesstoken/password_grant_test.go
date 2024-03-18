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

func TestServiceAccessToken_PasswordGrant(t *testing.T) {
	Convey("Test Password Grant", t, func() {
		// ctx
		ctx := context.Background()

		// mock
		out := &builder.MAccessTokenBuilder{}
		uc := &usecase.MAccessTokenUsecase{}

		Convey("Positive Scenario", func() {
			// set request
			req := &pb.PasswordGrantRequest{
				ClientId:     "TEST",
				ClientSecret: "TEST",
				Username:     "TEST",
				Password:     "TEST",
				Scopes:       "TEST",
			}

			// builder response
			resBuilder := &pb.AccessTokenResponse{
				Code:         0,
				Message:      "TEST",
				ErrorMessage: "TEST",
				Data:         nil,
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

			// mock builder
			out.On("AccessTokenResponse", mock.Anything, mock.Anything, mock.Anything).Return(resBuilder, nil).Once()

			// mock usecase
			uc.On("PasswordGrant", mock.Anything, mock.Anything).Return(resUC, nil).Once()

			// new service
			svc := NewAccessTokenService(out, uc)

			// password grant
			res, err := svc.PasswordGrant(ctx, req)

			// error
			Convey("error checking", func() {
				So(err, ShouldBeNil)
			})

			// response
			Convey("response checking", func() {
				So(res, ShouldNotBeNil)
			})
		})

		Convey("Negative Scenario", func() {
			// set request
			req := &pb.PasswordGrantRequest{
				ClientId:     "TEST",
				ClientSecret: "TEST",
				Username:     "TEST",
				Password:     "TEST",
				Scopes:       "TEST",
			}

			// error
			errRes := errors.New("TEST")

			// mock usecase
			uc.On("PasswordGrant", mock.Anything, mock.Anything).Return(nil, errRes).Once()

			// new service
			svc := NewAccessTokenService(out, uc)

			// password grant
			res, err := svc.PasswordGrant(ctx, req)

			// error
			Convey("error checking", func() {
				So(err, ShouldNotBeNil)
			})

			// response
			Convey("response checking", func() {
				So(res, ShouldBeNil)
			})
		})
	})
}
