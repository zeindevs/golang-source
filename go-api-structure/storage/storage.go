package storage

import "github.com/zeindevs/go-api-structure/types"

type Storage interface {
	Get(int) *types.User
}
