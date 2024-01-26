//go:build wireinject
// +build wireinject

package container

import (
	"github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	register "github.com/evenyosua18/auth2/app/usecase/registration"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeRegistrationUsecase(connection *db.MongoConnection) (registration register.IRegistrationUsecase) {
	wire.Build(user.NewUserRepository, accesstoken.NewAccessTokenRepository, refreshtoken.NewRefreshTokenRepository, register.NewRegistrationUsecase)
	return
}
