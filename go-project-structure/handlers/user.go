package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zeindevs/go-project-structure/data"
)

type Account struct {
	ID              int    `json:"id"`
	OnboardingState int    `json:"onboardingState"`
	AccountType     string `json:"accountType"`
}

type UserWithAccount struct {
	ID      uuid.UUID `json:"id"`
	Account Account   `json:"account"`
}

func HandleGetUser(c echo.Context) error {
	user := getAuthenticatedUser(c)
	resp := UserWithAccount{
		ID: user.ID,
	}
	userDetails, err := data.GetUserDetails(user.ID)
	if err != nil {
		return c.JSON(http.StatusOK, resp)
	}
	resp.Account = Account{
		ID:              int(userDetails.AccountID),
		OnboardingState: int(userDetails.Account.OnboardingState),
		AccountType:     string(userDetails.Account.AccountType),
	}
	return c.JSON(http.StatusOK, resp)
}
