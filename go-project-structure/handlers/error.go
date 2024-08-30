package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return e.Message
}

func Error(status int, msg string) APIError {
	return APIError{
		Status:  status,
		Message: msg,
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	slog.Error("api error", "err", err)
	apiError, ok := err.(APIError)
	if ok {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]string{"msg": "internal server error"})
}
