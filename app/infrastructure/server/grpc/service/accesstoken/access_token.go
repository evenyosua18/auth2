package accesstoken

import (
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/usecase/accesstoken"
)

type ServiceAccessToken struct {
	uc accesstoken.IAccessTokenUsecase
	pb.UnimplementedAccessTokenServiceServer
}

func NewAccessTokenService(uc accesstoken.IAccessTokenUsecase) *ServiceAccessToken {
	return &ServiceAccessToken{uc: uc}
}
