package registration

import (
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRegistrationUsecase(t *testing.T) {
	Convey("Test Registration Usecase Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock repository
			userRepo := &authdb.MUserRepository{}
			accessTokenRepo := &authdb.MAccessTokenRepository{}
			refreshTokenRepo := &authdb.MRefreshTokenRepository{}

			res := NewRegistrationUsecase(userRepo, accessTokenRepo, refreshTokenRepo)

			// response should be valid
			Convey("validate response should be valid", func() {
				So(res, ShouldEqual, &UsecaseRegistration{
					user:         userRepo,
					accessToken:  accessTokenRepo,
					refreshToken: refreshTokenRepo,
				})
			})
		})
	})
}
