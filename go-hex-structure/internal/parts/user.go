package parts

import (
	"context"
	"errors"
	"hex-structure/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type UserRepo interface {
  Add(ctx context.Context, u user.User) error
}
