package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/mitchellh/mapstructure"
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

	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	user := userRes.(*model.UserModel)

	// check user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 402))
	}

	res, err := u.manageAccessToken(sp, user, 1)

	if err != nil {
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
