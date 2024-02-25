package registration

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

func TestServiceRegistration_RegisterUser(t *testing.T) {
	Convey("Test Register User", t, func() {
		// context
		ctx := context.Background()

		// mock
		out := &builder.MRegistrationBuilder{}
		uc := &usecase.MRegistrationUsecase{}

		Convey("Positive Scenario", func() {
			Convey("all functions works", func() {
				// set request
				req := &pb.RegistrationUserRequest{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// set response usecase
				resUC := struct {
					AccessToken  string
					RefreshToken string
					ExpiredAt    int64
					Id           string
				}{
					AccessToken:  "TEST",
					RefreshToken: "TEST",
					ExpiredAt:    0,
					Id:           "TEST",
				}

				// set response builder
				resBuilder := pb.RegistrationUserResponse{
					Code:         201,
					Message:      "",
					ErrorMessage: "",
					Data:         nil,
				}

				// mock builder
				out.On("RegistrationUserResponse", mock.Anything, mock.Anything, mock.Anything).Return(&resBuilder, nil).Once()

				// mock usecase
				uc.On("RegistrationUser", mock.Anything, mock.Anything).Return(&resUC, nil).Once()

				// create service
				svc := NewRegistrationService(out, uc)

				// register user
				res, err := svc.RegisterUser(ctx, req)

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
			Convey("error when register user usecase", func() {
				// set request
				req := &pb.RegistrationUserRequest{
					Username: "TEST",
					Password: "TEST",
					Email:    "TEST",
					Phone:    "TEST",
				}

				// error response
				errRes := errors.New("TEST")

				// mock usecase
				uc.On("RegistrationUser", mock.Anything, mock.Anything).Return(nil, errRes).Once()

				// create service
				svc := NewRegistrationService(out, uc)

				// register user
				res, err := svc.RegisterUser(ctx, req)

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
	})
}
