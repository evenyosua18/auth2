package registration

import (
	"github.com/evenyosua18/auth2/app/mock/builder"
	"github.com/evenyosua18/auth2/app/mock/usecase"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRegistrationService(t *testing.T) {
	Convey("Test Registration Service Creation", t, func() {
		Convey("creation should return valid struct", func() {
			// mock builder
			registrationBuilder := &builder.MRegistrationBuilder{}

			// mock usecase
			registrationUC := &usecase.MRegistrationUsecase{}

			// new service
			res := NewRegistrationService(registrationBuilder, registrationUC)

			// result
			Convey("result should be valid", func() {
				So(res, ShouldEqual, &ServiceRegistration{
					uc:  registrationUC,
					out: registrationBuilder,
				})
			})
		})
	})
}
