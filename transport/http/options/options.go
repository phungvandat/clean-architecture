package options

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"github.com/phungvandat/clean-architecture/util/config/env"
	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/errors"
)

func LogRequestInfo(logger log.Logger) func(ctx context.Context, req *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		logger.Log("Method", req.Method, "Route", req.RequestURI)
		return ctx
	}
}

func VerifyToken(ctx context.Context, req *http.Request) context.Context {
	accessToken := req.Header.Get("Authorization")
	if strings.Trim(accessToken, " ") != "" {
		claims, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.InvalidSigningAlgorithm
			}
			secret := env.GetJWTSerectKeyEnv()
			return []byte(secret), nil
		})

		if err != nil || !claims.Valid {
			goto End
		}
		data := claims.Claims.(jwt.MapClaims)
		userID, check := data["user_id"].(string)

		if check {
			ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)
		}

		username, check := data["username"].(string)
		if check {
			ctx = context.WithValue(ctx, constants.UsernameContextKey, username)
		}

		role, check := data["role"].(string)
		if check {
			ctx = context.WithValue(ctx, constants.UserRoleContextKey, role)
		}
	}
End:
	return ctx
}
