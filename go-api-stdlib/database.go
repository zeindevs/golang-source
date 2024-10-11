package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "user1@gmail.com",
		Id:    "user1",
		Name:  "User 1",
		Token: "123",
	},
	"user2": {
		Email: "user2@gmail.com",
		Id:    "user2",
		Name:  "User 2",
		Token: "123",
	},
}
