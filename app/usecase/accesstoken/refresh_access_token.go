package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/utils/str"
	"github.com/evenyosua18/auth2/app/utils/token"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"strconv"
	"time"
)

type RefreshAccessTokenRequest struct {
	AccessToken  string
	RefreshToken string
}

func (u *UsecaseAccessToken) RefreshAccessToken(ctx context.Context, in interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode refresh token
	var req RefreshAccessTokenRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// get claims
	claims, err := token.ValidateToken(tracing.Context(sp), req.AccessToken)
	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// check token expired
	if claims[constant.ClaimsExpired].(int64) > time.Now().Unix() {
		return nil, tracing.LogError(sp, codes.Wrap(err, 427))
	}

	// get refresh token & access token
	refreshTokenRes, err := u.refreshToken.GetRefreshToken(ctx, []bson.M{
		{
			"$match": bson.M{
				"deleted_at":    nil,
				"refresh_token": req.RefreshToken,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "access_tokens",
				"localField":   "access_token_id",
				"foreignField": "_id",
				"as":           "access_token",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$access_token",
				"preserveNullAndEmptyArrays": true,
			},
		},
	})

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	// decode refresh token
	refreshToken, ok := refreshTokenRes.(model.RefreshTokenModel)

	if !ok {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	// check max refresh token
	if maxRefreshToken, err := strconv.Atoi(os.Getenv(constant.MaxRefreshToken)); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	} else if maxRefreshToken <= refreshToken.Count {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 428))
	}

	// validate access token id
	if refreshToken.AccessTokenId.Hex() != claims[constant.ClaimsId].(string) {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 426))
	}

	// generate access token
	tokenStr, expiredAt, err := token.GenerateToken(tracing.Context(sp), refreshToken.AccessTokenId.Hex(), token.ClaimsInformation{
		Username: claims[constant.ClaimsUsername].(string),
		Phone:    claims[constant.ClaimsPhone].(string),
		Email:    claims[constant.ClaimsEmail].(string),
	})

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 503))
	}

	// manage refresh token
	newRefreshToken, err := u.generateNewRefreshToken(sp, refreshToken)

	if err != nil {
		return nil, err
	}

	// update access token
	if _, err := u.accessToken.UpdateAccessToken(tracing.Context(sp), bson.M{
		"_id":        refreshToken.AccessTokenId,
		"deleted_at": nil,
	}, bson.M{
		"expired_at": expiredAt,
		"updated_at": time.Now(),
	}); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	return u.out.AccessTokenResponse(tracing.Context(sp), struct {
		AccessToken  string
		RefreshToken string
		ExpireAt     int64
	}{
		AccessToken:  tokenStr,
		RefreshToken: newRefreshToken.RefreshToken,
		ExpireAt:     expiredAt.Unix(),
	}, 200)
}

func (u *UsecaseAccessToken) generateNewRefreshToken(sp interface{}, prevRefreshToken model.RefreshTokenModel) (refreshToken *model.RefreshTokenModel, err error) {
	// generate new refresh token
	refreshToken = &model.RefreshTokenModel{
		Id:            primitive.NewObjectID(),
		AccessTokenId: prevRefreshToken.AccessTokenId,
		RefreshToken:  str.GenerateString(16, ""),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Count:         prevRefreshToken.Count + 1,
		UserId:        prevRefreshToken.UserId,
	}

	// saved new refresh token
	if err = u.refreshToken.InsertRefreshToken(tracing.Context(sp), refreshToken); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// delete refresh token
	if err = u.refreshToken.DeleteRefreshToken(tracing.Context(sp), bson.M{
		"_id":        prevRefreshToken.Id,
		"deleted_at": nil,
	}); err != nil && err != mongo.ErrNoDocuments {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	}

	return refreshToken, nil
}