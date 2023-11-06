// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/builder"
	accesstoken2 "github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	"github.com/evenyosua18/auth2/app/usecase/accesstoken"
	oauthclient2 "github.com/evenyosua18/auth2/app/usecase/oauthclient"
	"github.com/evenyosua18/auth2/app/usecase/registration"
	"github.com/evenyosua18/auth2/app/utils/db"
)

// Injectors from access_token_usecase.go:

func InitializeAccessTokenUsecase(connection *db.MongoConnection) accesstoken.IAccessTokenUsecase {
	iAccessTokenBuilder := builder.NewAccessTokenBuilder()
	iOauthClientRepository := oauthclient.NewOauthClientRepository(connection)
	iUserRepository := user.NewUserRepository(connection)
	iAccessTokenRepository := accesstoken2.NewAccessTokenRepository(connection)
	iRefreshTokenRepository := refreshtoken.NewRefreshTokenRepository(connection)
	iAccessTokenUsecase := accesstoken.NewAccessTokenUsecase(iAccessTokenBuilder, iOauthClientRepository, iUserRepository, iAccessTokenRepository, iRefreshTokenRepository)
	return iAccessTokenUsecase
}

// Injectors from oauth_client_usecase.go:

func InitializeOauthClientUsecase(connection *db.MongoConnection) oauthclient2.IOauthClientUsecase {
	iOauthClientRepository := oauthclient.NewOauthClientRepository(connection)
	iOauthClientUsecase := oauthclient2.NewOauthClientUsecase(iOauthClientRepository)
	return iOauthClientUsecase
}

// Injectors from registration_usecase.go:

func InitializeRegistrationUsecase(connection *db.MongoConnection) registration.IRegistrationUsecase {
	iRegistrationBuilder := builder.NewRegistrationBuilder()
	iUserRepository := user.NewUserRepository(connection)
	iAccessTokenRepository := accesstoken2.NewAccessTokenRepository(connection)
	iRefreshTokenRepository := refreshtoken.NewRefreshTokenRepository(connection)
	iRegistrationUsecase := registration.NewRegistrationUsecase(iRegistrationBuilder, iUserRepository, iAccessTokenRepository, iRefreshTokenRepository)
	return iRegistrationUsecase
}
