package oauthclient

import (
	"github.com/evenyosua18/auth2/app/mock/repository/authdb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewOauthClientUsecase(t *testing.T) {
	Convey("Test Oauth Client Usecase Creation", t, func() {
		// mock oauth client repository
		oauthClientRepo := &authdb.MOauthClientRepository{}

		// create oauth client usecase
		res := NewOauthClientUsecase(oauthClientRepo)

		// response should be valid
		Convey("validate response should be valid", func() {
			So(res, ShouldEqual, &UsecaseOauthClient{
				oauthClient: oauthClientRepo,
			})
		})
	})
}
