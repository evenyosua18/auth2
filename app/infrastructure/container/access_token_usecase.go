//go:build wireinject
// +build wireinject

package container

import (
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/builder"
	accessTokenRepo "github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	"github.com/evenyosua18/auth2/app/usecase/accesstoken"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeAccessTokenUsecase(connection *db.MongoConnection) (accessToken accesstoken.IAccessTokenUsecase) {
	wire.Build(user.NewUserRepository, oauthclient.NewOauthClientRepository, accessTokenRepo.NewAccessTokenRepository, refreshtoken.NewRefreshTokenRepository, builder.NewAccessTokenBuilder, accesstoken.NewAccessTokenUsecase)
	return
}
