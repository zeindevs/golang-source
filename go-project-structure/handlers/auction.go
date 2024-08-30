package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zeindevs/go-project-structure/types"
)

func HandleGetAuctions(c echo.Context) error {
	user := getAuthenticatedUser(c)
	if user.OnboardingState == types.OnboardingStateBasic {

	}
	return nil
 }

func getDemoAuctions() {

}
