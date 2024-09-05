package main

import (
	"fmt"
	"log"
)

type User struct {
	id    int
	email string
	age   int
}

func InsertUser(user *User) error {
	user.id = 1
	return nil
}

func GetUserByID(id int) (User, error) {
	return User{}, nil
}

func main() {
	user := User{
    age: 20,
    email: "user@mail.com",
  }
  if err := InsertUser(&user); err != nil {
    log.Fatal(err)
  }

	fmt.Println(user)
}
