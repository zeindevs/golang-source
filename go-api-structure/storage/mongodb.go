package storage

import "github.com/zeindevs/go-api-structure/types"

type MongoStorage struct{}

func NewMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "foo",
	}
}
