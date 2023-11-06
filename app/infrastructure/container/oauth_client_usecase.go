//go:build wireinject
// +build wireinject

package container

import (
	oauthClientRepo "github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
	"github.com/evenyosua18/auth2/app/usecase/oauthclient"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/google/wire"
)

func InitializeOauthClientUsecase(connection *db.MongoConnection) (oauthClient oauthclient.IOauthClientUsecase) {
	wire.Build(oauthClientRepo.NewOauthClientRepository, oauthclient.NewOauthClientUsecase)
	return
}
