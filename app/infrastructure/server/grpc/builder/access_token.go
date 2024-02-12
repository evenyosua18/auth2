package builder

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/accesstoken"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
)

type AccessTokenBuilder struct{}

func NewAccessTokenBuilder() accesstoken.IAccessTokenBuilder {
	return &AccessTokenBuilder{}
}

func (b *AccessTokenBuilder) AccessTokenResponse(ctx context.Context, in interface{}, code int) (*pb.AccessTokenResponse, error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	//decode response to proto format
	var data *pb.AccessTokenData
	if err := mapstructure.Decode(in, &data); err != nil {
		return nil, err
	}

	// get custom code for response
	c := codes.Get(code)

	// manage response
	res := &pb.AccessTokenResponse{
		Code:         int32(code),
		Message:      c.ResponseMessage,
		ErrorMessage: c.ErrorMessage,
		Data:         data,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
