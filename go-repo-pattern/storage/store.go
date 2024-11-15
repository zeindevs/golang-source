package storage

type User struct {
	ID    uint
	Name  string
	Email string
}

type Storage interface {
	CreateUser(user *User) error
}
