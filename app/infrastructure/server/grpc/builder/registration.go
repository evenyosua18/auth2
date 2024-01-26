package builder

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/registration"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
)

type RegistrationBuilder struct{}

func NewRegistrationBuilder() registration.IRegistrationBuilder {
	return &RegistrationBuilder{}
}

func (r *RegistrationBuilder) RegistrationUserResponse(ctx context.Context, in interface{}, code int) (*pb.RegistrationUserResponse, error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode data to proto format
	var data *pb.RegistrationUserData
	if err := mapstructure.Decode(in, &data); err != nil {
		return nil, err
	}

	// get custom code for response
	c := codes.Get(code)

	// manage response
	res := &pb.RegistrationUserResponse{
		Code:         int32(code),
		Message:      c.ResponseMessage,
		ErrorMessage: c.ErrorMessage,
		Data:         data,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
