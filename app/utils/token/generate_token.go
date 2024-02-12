package token

import (
	"context"
	"crypto/rsa"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/tracing"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type ClaimsInformation struct {
	Username string
	Phone    string
	Email    string
}

func GenerateToken(ctx context.Context, uuid string, info ClaimsInformation) (token string, expiredAt time.Time, err error) {
	// start span
	sp := tracing.StartChild(ctx, uuid, info)
	defer tracing.Close(sp)

	// generate token duration
	var maxAge time.Duration
	maxAge, err = time.ParseDuration(os.Getenv(constant.TokenDuration) + "h")

	if err != nil {
		tracing.LogError(sp, err)
		return
	}

	// determine expired at
	expiredAt = time.Now().Add(maxAge)
	tracing.LogObject(sp, "token_expired_at", expiredAt)

	// read rsa key
	var rsaKey []byte
	rsaKey, err = os.ReadFile(os.Getenv(constant.TokenPrivateKey))

	if err != nil {
		tracing.LogError(sp, err)
		return
	}

	var key *rsa.PrivateKey
	key, err = jwt.ParseRSAPrivateKeyFromPEM(rsaKey)
	if err != nil {
		tracing.LogError(sp, err)
		return
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		constant.ClaimsId:       uuid,
		constant.ClaimsUsername: info.Username,
		constant.ClaimsEmail:    info.Email,
		constant.ClaimsPhone:    info.Phone,
		constant.ClaimsExpired:  expiredAt.Unix(),
	}).SignedString(key)

	if err != nil {
		tracing.LogError(sp, err)
		return
	}

	return
}
