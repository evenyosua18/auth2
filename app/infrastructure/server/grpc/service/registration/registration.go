package registration

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/usecase/registration"
)

type IRegistrationBuilder interface {
	RegistrationUserResponse(ctx context.Context, in interface{}, code int) (*pb.RegistrationUserResponse, error)
}

type ServiceRegistration struct {
	uc  registration.IRegistrationUsecase
	out IRegistrationBuilder
	pb.UnimplementedRegistrationServiceServer
}

func NewRegistrationService(out IRegistrationBuilder, uc registration.IRegistrationUsecase) *ServiceRegistration {
	return &ServiceRegistration{
		out: out,
		uc:  uc,
	}
}
