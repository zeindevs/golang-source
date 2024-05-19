package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUnAuthenticated = errors.New("unauthenticated")
	ErrTokenInvalid    = errors.New("token invalid")
)

type AdminAuthMiddleware struct{}

func (am *AdminAuthMiddleware) Authenticate(w http.ResponseWriter, r *http.Request) error {
	tokenStr := r.Header.Get("x-api-token")
	if len(tokenStr) == 0 {
		return ErrUnAuthenticated
	}

	token, err := parseJWT(tokenStr)
	if err != nil {
		return ErrUnAuthenticated
	}

	if !token.Valid {
		return ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrUnAuthenticated
	}

	id := claims["userID"]
	_ = id

	fmt.Println("guarding the admin routes")
	return nil
}

func parseJWT(tokenStr string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}
