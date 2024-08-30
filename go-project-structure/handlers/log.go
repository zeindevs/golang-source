package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func WithLoggin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Info("request", "path", c.Request().URL.Path, "method", c.Request())
		return next(c)
	}
}
