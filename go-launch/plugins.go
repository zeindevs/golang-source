package launch

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

func errorHandlerPlug(next Handler) Handler {
	_ = sync.Pool{}
	return func(c *Ctx) error {
		c.ErrorHandler = func(c *Ctx, err error) {
			if launchErr, ok := err.(Error); ok {
				c.Status(launchErr.StatusCode).JSON(launchErr)
				return
			}
			c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}

func requestIdPlug(next Handler) Handler {
	return func(c *Ctx) error {
		c.RequestID = uuid.Must(uuid.NewV7()).String()
		return next(c)
	}
}

func notFoundPlugin(next Handler) Handler {
	return func(c *Ctx) error {
		c.launch.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.View("errors/404", nil)
			return
		})
		return nil
	}
}
