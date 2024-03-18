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

func TestServiceAccessToken_ValidateAccessToken(t *testing.T) {
	Convey("Test Validate Access Token", t, func() {
		// context
		ctx := context.Background()

		// mock
		out := &builder.MAccessTokenBuilder{}
		uc := &usecase.MAccessTokenUsecase{}

		Convey("Positive Scenario", func() {
			// set request
			req := &pb.ValidateTokenRequest{AccessToken: "TEST"}

			// mock usecase
			uc.On("ValidateAccessToken", mock.Anything, mock.Anything).Return(nil).Once()

			// new service
			svc := NewAccessTokenService(out, uc)

			// validate access token
			res, err := svc.ValidateAccessToken(ctx, req)

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
			req := &pb.ValidateTokenRequest{AccessToken: "TEST"}

			// error
			errRes := errors.New("TEST")

			// mock usecase
			uc.On("ValidateAccessToken", mock.Anything, mock.Anything).Return(errRes).Once()

			// new service
			svc := NewAccessTokenService(out, uc)

			// validate access token
			res, err := svc.ValidateAccessToken(ctx, req)

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
