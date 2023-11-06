package registration

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/ego-util/tracing"
	"google.golang.org/grpc/status"
)

func (s *ServiceRegistration) RegisterUser(ctx context.Context, in *pb.RegistrationUserRequest) (*pb.RegistrationUserResponse, error) {
	//trace
	sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, "RegisterUser"))
	defer tracing.Close(sp)

	//call interaction
	res, err := s.uc.RegistrationUser(tracing.Context(sp), in)
	if err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	tracing.LogResponse(sp, res)
	return res.(*pb.RegistrationUserResponse), nil
}
