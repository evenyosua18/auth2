package accesstoken

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type PasswordGrantRequest struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	Scopes       string
}

func (u *UsecaseAccessToken) PasswordGrant(ctx context.Context, in interface{}) (interface{}, error) {
	// tracer
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode request
	var req PasswordGrantRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}
	tracing.LogObject(sp, "request after decode", req)

	//get user
	userRes, err := u.user.GetUser(tracing.Context(sp), struct {
		Username string
	}{
		Username: req.Username,
	})

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, tracing.LogError(sp, codes.Wrap(err, 403))
	} else if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 501))
	}

	user := userRes.(*model.UserModel)

	// check user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 402))
	}
	/*
		// create token model
		savedToken := model.AccessTokenModel{
			Id:        primitive.NewObjectID(),
			UserId:    user.Id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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
			Id:            primitive.NewObjectID(),
			AccessTokenId: savedToken.Id,
			RefreshToken:  str.GenerateString(16, ""),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			Count:         1,
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
	*/

	res, err := u.manageAccessToken(sp, user, 1)

	if err != nil {
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
