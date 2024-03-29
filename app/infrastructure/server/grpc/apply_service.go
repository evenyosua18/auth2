package grpc

import (
	"github.com/evenyosua18/auth2/app/infrastructure/container"
	"github.com/evenyosua18/auth2/app/infrastructure/proto/pb"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/builder"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/accesstoken"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/registration"
	"github.com/evenyosua18/auth2/app/repository"
	"google.golang.org/grpc"
)

// Apply register all service here
func Apply(server *grpc.Server) {
	pb.RegisterAccessTokenServiceServer(server, accesstoken.NewAccessTokenService(builder.NewAccessTokenBuilder(), container.InitializeAccessTokenUsecase(repository.Con.MainMongoDB)))
	pb.RegisterRegistrationServiceServer(server, registration.NewRegistrationService(builder.NewRegistrationBuilder(), container.InitializeRegistrationUserUsecase(repository.Con.MainMongoDB)))
	//pb.RegisterUserServiceServer(server, user.NewUserService(container.InitializeUserInteraction()))
}
