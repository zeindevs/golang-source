package main

import "fmt"

// When
// 1. When we need to update state
// pointer = 8 bytes
// 2. When we need to optimize the memory large objects that are getting called A LOT.

type User struct {
	email    string
	username string
	age      int
	file     []byte // ?? Large ??
}

func getUser() (*User, error) {
	// return User{}, fmt.Errorf("foo")
	return nil, fmt.Errorf("foo")
}

func (u User) Email() string {
	return u.email
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

// x amount of bytes => sizeOf(user)
// 1 gb user size
func Email(user User) string {
	return user.email
}

func main() {
	user := User{
		email: "agg@foo.com",
	}
	user.UpdateEmail("foo@gmail.com")
	fmt.Println(user.Email())
}
