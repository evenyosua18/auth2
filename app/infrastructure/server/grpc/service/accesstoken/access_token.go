package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/usecase/accesstoken"
)

type IAccessTokenBuilder interface {
	AccessTokenResponse(ctx context.Context, in interface{}, code int) (*pb.AccessTokenResponse, error)
}

type ServiceAccessToken struct {
	uc  accesstoken.IAccessTokenUsecase
	out IAccessTokenBuilder
	pb.UnimplementedAccessTokenServiceServer
}

func NewAccessTokenService(out IAccessTokenBuilder, uc accesstoken.IAccessTokenUsecase) *ServiceAccessToken {
	return &ServiceAccessToken{uc: uc, out: out}
}
