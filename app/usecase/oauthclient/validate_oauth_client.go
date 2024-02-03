package oauthclient

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type ValidateOauthClientRequest struct {
	ClientId     string
	ClientSecret string
}

func (u *UsecaseOauthClient) ValidateOauthClient(ctx context.Context, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode request
	var req ValidateOauthClientRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// get oauth client
	oauthClientRes, err := u.oauthClient.GetOauthClient(tracing.Context(sp), struct {
		ClientId string
	}{
		ClientId: req.ClientId,
	})

	if err != nil {
		return tracing.LogError(sp, err)
	} else if oauthClientRes == nil {
		return tracing.LogError(sp, codes.Wrap(err, 420))
	}

	oauthClient, ok := oauthClientRes.(*model.OauthClientModel)

	if !ok {
		return tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// check client id
	if oauthClient.ClientId != req.ClientId {
		return tracing.LogError(sp, codes.Wrap(nil, 421))
	}

	// compare client secret
	if err := bcrypt.CompareHashAndPassword([]byte(oauthClient.ClientSecret), []byte(req.ClientSecret)); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 421))
	}

	return nil
}
