package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zeindevs/go-project-structure/data"
	"github.com/zeindevs/go-project-structure/types"
)

type UserMetadata struct {
	AccountType string `json:"accountType"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	UserMetadata UserMetadata `json:"user_metadata"`
	Email        string       `json:"email"`
	ID           uuid.UUID    `json:"sub"`
}

func WithAccountTypeOnly(accountType types.AccountType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := getAuthenticatedUser(c)
			if user.AccountID == 0 {
				return Error(http.StatusUnauthorized, "unauthorized")
			}
			if user.AccountType != accountType {
				return Error(http.StatusUnauthorized, "unauthorized")
			}
			return next(c)
		}
	}
}

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if len(token) == 0 {
			fmt.Println("Authorization token not set")
			return Error(http.StatusUnauthorized, "unauthorized")
		}
		claims, valid := validateJWT(token)
		if !valid {
			fmt.Println("invalid JWT")
			return Error(http.StatusUnauthorized, "unauthorized")
		}
		authUser := &types.AuthenticatedUser{
			ID:    claims.ID,
			Email: claims.Email,
		}
		userDetails, err := data.GetUserDetails(claims.ID)
		if err == nil && userDetails != nil {
			authUser.OnboardingState = userDetails.Account.OnboardingState
			authUser.AccountType = userDetails.Account.AccountType
			authUser.AccountID = int32(userDetails.AccountID)
		}
		c.Set("authUser", authUser)
		return next(c)
	}
}

func validateJWT(tokenString string) (*JWTClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return nil, nil
	})
	if err != nil {
		return nil, false
	}
	_ = token
	return &JWTClaims{}, true
}

func getAuthenticatedUser(c echo.Context) *types.AuthenticatedUser {
	return &types.AuthenticatedUser{}
}
