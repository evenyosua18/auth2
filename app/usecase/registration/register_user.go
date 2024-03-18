package registration

import (
	"context"
	"errors"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/repository"
	"github.com/evenyosua18/auth2/app/utils/str"
	"github.com/evenyosua18/auth2/app/utils/token"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type UserRegistrationRequest struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type UserRegistrationResponse struct {
	Id           string
	RefreshToken string
	AccessToken  string
	ExpiredAt    int64
}

func (u *UsecaseRegistration) RegistrationUser(ctx context.Context, in interface{}) (interface{}, error) {
	// tracing
	sp := tracing.StartChild(ctx, in)
	defer tracing.Close(sp)

	// decode request
	var req UserRegistrationRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 502))
	}

	// check username, email, phone exist
	if user, err := u.user.GetUser(tracing.Context(sp), struct {
		Username   string
		Email      string
		Phone      string
		CheckExist bool
	}{
		Username:   req.Username,
		Email:      req.Email,
		Phone:      req.Phone,
		CheckExist: true,
	}); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if user != nil {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 410))
	}

	// encrypt user password
	userPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, tracing.LogError(sp, codes.Wrap(err, 503))
	}

	// create user model
	user := model.UserModel{
		Id:       primitive.NewObjectID(),
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
		Password: string(userPassword),
		IsActive: true,
	}

	// create token model
	savedToken := model.AccessTokenModel{
		Id:     primitive.NewObjectID(),
		UserId: user.Id,
	}

	// generate access token
	tokenStr, expiredAt, err := token.GenerateToken(tracing.Context(sp), savedToken.Id.Hex(), token.ClaimsInformation{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
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
		Count:         1,
		UserId:        user.Id,
	}

	// set ctx
	childCtx := tracing.Context(sp)

	// start session
	if os.Getenv(constant.UseReplica) == constant.True {
		if err := repository.Con.MainMongoDB.StartSession(tracing.GetTraceID(sp)); err != nil {
			return nil, tracing.LogError(sp, codes.Wrap(err, 504))
		}

		// setup trace id to ctx
		childCtx = tracing.AddContextValue(sp, constant.SessionId, tracing.GetTraceID(sp))
	}

	// save user
	if err := u.user.InsertUser(childCtx, user); err != nil {
		return nil, tracing.LogError(sp, err)
	}

	// save access token
	if err := u.accessToken.InsertAccessToken(childCtx, savedToken); err != nil {
		return nil, tracing.LogError(sp, err)
	}

	// save refresh token
	if err := u.refreshToken.InsertRefreshToken(childCtx, refreshToken.SetCreated()); err != nil {
		return nil, tracing.LogError(sp, err)
	}

	// commit session
	if os.Getenv(constant.UseReplica) == constant.True {
		repository.Con.MainMongoDB.EndSession(childCtx, tracing.GetTraceID(sp))
	}

	tracing.LogResponse(sp, savedToken)
	return UserRegistrationResponse{
		Id:           savedToken.Id.Hex(),
		RefreshToken: refreshToken.RefreshToken,
		AccessToken:  tokenStr,
		ExpiredAt:    expiredAt.Unix(),
	}, nil
}
