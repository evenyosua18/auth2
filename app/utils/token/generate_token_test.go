package token

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	Convey("Test Generate Token Utility", t, func() {
		// context
		ctx := context.Background()

		// set env
		os.Setenv(constant.TokenDuration, "24")
		os.Setenv(constant.TokenSignature, "Oauth2 Testing")

		// set request data
		uuid := primitive.NewObjectID()
		claimsRequest := ClaimsInformation{
			Username: "TEST",
			Phone:    "TEST",
			Email:    "TEST",
		}

		// generate token
		tok, expiredAt, err := GenerateToken(ctx, uuid.Hex(), claimsRequest)

		// error
		Convey("error checking", func() {
			So(err, ShouldBeNil)
		})

		// expired at
		Convey("expire checking", func() {
			So(expiredAt.Unix(), ShouldBeGreaterThan, time.Now().Add(time.Minute*-1).Unix())
		})

		// token
		Convey("token checking", func() {
			So(t, ShouldNotBeEmpty)

			// claims
			Convey("extract claims from token", func() {
				claims, err := ValidateToken(ctx, tok)

				So(err, ShouldBeNil)
				So(claims[constant.ClaimsExpired], ShouldHaveSameTypeAs, expiredAt.Unix())
				So(expiredAt.Unix(), ShouldEqual, claims[constant.ClaimsExpired])
				So(claims[constant.ClaimsUsername], ShouldEqual, claimsRequest.Username)
				So(claims[constant.ClaimsPhone], ShouldEqual, claimsRequest.Phone)
				So(claims[constant.ClaimsEmail], ShouldEqual, claimsRequest.Email)
				So(claims[constant.ClaimsId], ShouldEqual, uuid.Hex())
			})
		})
	})
}
