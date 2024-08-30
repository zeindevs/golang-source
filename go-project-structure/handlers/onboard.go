package handlers

import (
  "encoding/json"
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/zeindevs/go-project-structure/data"
)

func HandlePostOnboard(c echo.Context) error {
  user := getAuthenticatedUser(c)
  var params data.OnboardingParams
  if err := json.NewDecoder(c.Request().Body).Decode(&params); err != nil {
    return err
  }
  if err := data.OnboardUser(user.ID, params); err != nil {
    return err
  }
  return c.JSON(http.StatusOK, "ok")
}
