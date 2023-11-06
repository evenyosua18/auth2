package oauthclient

import (
	"context"
	"github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
)

type IOauthClientUsecase interface {
	ValidateOauthClient(ctx context.Context, in interface{}) error
}

type IOauthClientBuilder interface {
}

type UsecaseOauthClient struct {
	//out         IOauthClientBuilder
	oauthClient oauthclient.IOauthClientRepository
}

func NewOauthClientUsecase(oauthClientRepo oauthclient.IOauthClientRepository) IOauthClientUsecase {
	return &UsecaseOauthClient{
		oauthClient: oauthClientRepo,
	}
}
