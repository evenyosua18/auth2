package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/utils/token"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ValidateTokenRequest struct {
	AccessToken string
}

func (u *UsecaseAccessToken) ValidateAccessToken(ctx context.Context, in interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode request
	var req ValidateTokenRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// decode access token
	claims, err := token.ValidateToken(tracing.Context(sp), req.AccessToken)
	if err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// get claims id
	accessTokenId, ok := claims[constant.ClaimsId].(string)

	if !ok {
		return tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	accessTokenObjectId, err := primitive.ObjectIDFromHex(accessTokenId)

	if err != nil {
		return tracing.LogError(sp, codes.Wrap(err, 503))
	}

	// get access token from db
	accessTokenRes, err := u.accessToken.GetAccessToken(tracing.Context(sp), struct {
		Id *primitive.ObjectID
	}{
		Id: &accessTokenObjectId,
	})

	if err != nil {
		return tracing.LogError(sp, err)
	}

	// convert to model
	accessToken, ok := accessTokenRes.(*model.AccessTokenModel)

	if !ok {
		return tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// check expired at
	if accessToken.ExpiredAt.Before(time.Now()) {
		return tracing.LogError(sp, codes.Wrap(nil, 423))
	}

	return nil
}
