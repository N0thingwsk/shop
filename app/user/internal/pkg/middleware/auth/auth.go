package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

var currentUserKey struct{}

type CurrentUser struct {
	UserID uint
}

func GenerateToken(secret string, userid uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userid,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				auths := strings.SplitN(tokenString, " ", 2)
				token, err := jwt.Parse(auths[0], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					if u, ok := claims["userid"]; ok {
						ctx = context.WithValue(ctx, "userid", u)
					}
				} else {
					return nil, errors.New("Token Invalid")
				}
			}
			return handler(ctx, req)
		}
	}
}
