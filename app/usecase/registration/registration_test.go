package registration

import (
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRegistrationUserUsecase(t *testing.T) {
	Convey("Test Registration User Usecase Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock repository
			userRepo := &authdb.MUserRepository{}
			accessTokenRepo := &authdb.MAccessTokenRepository{}
			refreshTokenRepo := &authdb.MRefreshTokenRepository{}

			res := NewRegistrationUserUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

			// response should be valid
			Convey("validate response should be valid", func() {
				So(res, ShouldEqual, &UsecaseRegistrationUser{
					user:         userRepo,
					accessToken:  accessTokenRepo,
					refreshToken: refreshTokenRepo,
				})
			})
		})
	})
}

func TestNewRegistrationEndpointUsecase(t *testing.T) {
	Convey("Test Registration Endpoint Usecase Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock repository
			userRepo := &authdb.MUserRepository{}
			accessTokenRepo := &authdb.MAccessTokenRepository{}
			refreshTokenRepo := &authdb.MRefreshTokenRepository{}

			res := NewRegistrationUserUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

			// response should be valid
			Convey("validate response should be valid", func() {
				So(res, ShouldEqual, &UsecaseRegistrationUser{
					user:         userRepo,
					accessToken:  accessTokenRepo,
					refreshToken: refreshTokenRepo,
				})
			})
		})
	})
}
