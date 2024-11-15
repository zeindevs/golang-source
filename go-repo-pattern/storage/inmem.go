package storage

type InMemStorage struct {
	users []*User
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		users: []*User{},
	}
}

func (s *InMemStorage) CreateUser(user *User) error {
	s.users = append(s.users, user)
	return nil
}
