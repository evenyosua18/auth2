package accesstoken

import (
	"github.com/evenyosua18/auth2/app/mock/builder"
	"github.com/evenyosua18/auth2/app/mock/usecase"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewAccessTokenService(t *testing.T) {
	Convey("Test Access Token Service Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock builder
			accessTokenBuilder := &builder.MAccessTokenBuilder{}

			// mock usecase
			accessTokenUC := &usecase.MAccessTokenUsecase{}

			// new service
			res := NewAccessTokenService(accessTokenBuilder, accessTokenUC)

			// result
			Convey("result should be valid", func() {
				So(res, ShouldEqual, &ServiceAccessToken{
					uc:  accessTokenUC,
					out: accessTokenBuilder,
				})
			})
		})
	})
}
