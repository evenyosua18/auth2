package token

import (
	"context"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ValidateToken(ctx context.Context, token string) (jwt.MapClaims, error) {
	// tracing
	sp := tracing.StartChild(ctx)
	defer tracing.Close(sp)

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if jwtToken.Method.Alg() == "RS256" {
			//read rsa public key
			rsaPub, err := os.ReadFile(os.Getenv(constant.TokenPublicKey))
			if err != nil {
				return nil, tracing.LogError(sp, err)
			}

			key, err := jwt.ParseRSAPublicKeyFromPEM(rsaPub)
			if err != nil {
				return nil, tracing.LogError(sp, err)
			}

			return key, nil
		} else if jwtToken.Method.Alg() == "HS256" {
			return []byte(os.Getenv(constant.TokenSignature)), nil
		}

		return nil, tracing.LogError(sp, codes.Wrap(nil, 502))
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, tracing.LogError(sp, codes.Wrap(nil, 424))
	}

	claims[constant.ClaimsExpired] = int64(claims[constant.ClaimsExpired].(float64))

	tracing.LogResponse(sp, claims)
	return claims, nil
}
