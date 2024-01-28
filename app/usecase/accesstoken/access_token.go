package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/repository/authdb/accesstoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
	"github.com/evenyosua18/auth2/app/repository/authdb/refreshtoken"
	"github.com/evenyosua18/auth2/app/repository/authdb/user"
	"github.com/evenyosua18/auth2/app/utils/str"
	"github.com/evenyosua18/auth2/app/utils/token"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAccessTokenUsecase interface {
	PasswordGrant(ctx context.Context, in interface{}) (interface{}, error)
	ValidateAccessToken(ctx context.Context, in interface{}) error
	RefreshAccessToken(ctx context.Context, in interface{}) (interface{}, error)
}

type UsecaseAccessToken struct {
	oauthClient  oauthclient.IOauthClientRepository
	user         user.IUserRepository
	accessToken  accesstoken.IAccessTokenRepository
	refreshToken refreshtoken.IRefreshTokenRepository
}

func NewAccessTokenUsecase(oauthClientRepo oauthclient.IOauthClientRepository, userRepo user.IUserRepository, accessTokenRepo accesstoken.IAccessTokenRepository, refreshTokenRepo refreshtoken.IRefreshTokenRepository) IAccessTokenUsecase {
	return &UsecaseAccessToken{
		oauthClient:  oauthClientRepo,
		user:         userRepo,
		accessToken:  accessTokenRepo,
		refreshToken: refreshTokenRepo,
	}
}

// internal function only as an extension in use case layer, so still on same layer
// just for decrease redundancy

func (u *UsecaseAccessToken) manageAccessToken(sp interface{}, user *model.UserModel, count int) (interface{}, error) {
	// create token model
	savedToken := model.AccessTokenModel{
		Id:     primitive.NewObjectID(),
		UserId: user.Id,
	}

	// generate access token
	tokenStr, expiredAt, err := token.GenerateToken(tracing.Context(sp), savedToken.Id.Hex(), token.ClaimsInformation{
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
	})

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 503))
	}

	// update token expired at
	savedToken.ExpiredAt = expiredAt

	// create refresh token
	refreshToken := model.RefreshTokenModel{
		AccessTokenId: savedToken.Id,
		RefreshToken:  str.GenerateString(16, ""),
		Count:         count,
		UserId:        user.Id,
	}

	// save access token
	if err := u.accessToken.InsertAccessToken(tracing.Context(sp), savedToken); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// save refresh token
	if err := u.refreshToken.InsertRefreshToken(tracing.Context(sp), refreshToken); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// setup response
	return struct {
		AccessToken  string
		RefreshToken string
		ExpireAt     int64
	}{
		AccessToken:  tokenStr,
		RefreshToken: refreshToken.RefreshToken,
		ExpireAt:     savedToken.ExpiredAt.Unix(),
	}, nil
}
