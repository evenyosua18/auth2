package registration

import (
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/usecase/registration"
)

type ServiceRegistration struct {
	uc registration.IRegistrationUsecase
	pb.UnimplementedRegistrationServiceServer
}

func NewRegistrationService(uc registration.IRegistrationUsecase) *ServiceRegistration {
	return &ServiceRegistration{
		uc: uc,
	}
}
