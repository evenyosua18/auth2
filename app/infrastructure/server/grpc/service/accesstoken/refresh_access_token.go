package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/ego-util/tracing"
	"google.golang.org/grpc/status"
)

func (s *ServiceAccessToken) RefreshAccessToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.AccessTokenResponse, error) {
	//trace
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	//call interaction
	res, err := s.uc.RefreshAccessToken(tracing.Context(sp), in)
	if err != nil {
		return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
	}

	tracing.LogResponse(sp, res)
	return s.out.AccessTokenResponse(tracing.Context(sp), res, 200)
}
