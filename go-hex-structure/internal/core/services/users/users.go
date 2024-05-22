package users

import (
  "context"

  "hex-structure/internal/parts"
)

type API interface {
	CreateAccout(context.Context, CreateAccountReq) (*CreateAccountResp, error)
}

type Service struct {
	userRepo parts.UserRepo
}

func NewService(ur parts.UserRepo) *Service {
	return &Service{
		userRepo: ur,
	}
}
