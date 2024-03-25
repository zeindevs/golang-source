package data

import "github.com/zeindevs/launch/examples/handler/params"

type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}

func CreateUser(params params.CreateUser) *User {
	return &User{
		ID:                1,
		Email:             params.Email,
		EncryptedPassword: "faskafsafkasfk",
	}
}
