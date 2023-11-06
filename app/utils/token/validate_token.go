package token

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ValidateToken(ctx context.Context, token string) (jwt.MapClaims, error) {
	// tracing
	sp := tracing.StartChild(ctx)
	defer tracing.Close(sp)

	//read rsa public key
	rsaPub, err := os.ReadFile(os.Getenv(constant.TokenPublicKey))
	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(rsaPub)
	if err != nil {
		return nil, tracing.LogError(sp, err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 424))
	}

	tracing.LogResponse(sp, claims)
	return claims, nil
}
