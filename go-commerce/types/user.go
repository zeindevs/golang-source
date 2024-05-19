package types

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"_"`
	IsAdmin  bool   `bson:"isAdmin" json:"isAdmin"`
	Token    string `bson:"token" json:"token"`
}

func NewAdminUser(email, password string) (*User, error) {
	user, err := NewUser(email, password)
	if err != nil {
		return nil, err
	}

	user.IsAdmin = true

	return user, nil
}

func NewUser(email, password string) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Password: string(pass),
	}, nil
}

func (u *User) ValidatePassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}
