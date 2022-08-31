package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

var currentUserKey struct{}

type CurrentUser struct {
	UserID uint
}

// GenerateToken 生成JWTToken
func GenerateToken(secret string, userid uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userid,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
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
				fmt.Println(tokenString)
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}
