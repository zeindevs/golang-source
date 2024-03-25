package params

import "fmt"

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p CreateUser) Validate() error {
	if len(p.Email) != 4 {
		return fmt.Errorf("invalid email")
	}
	return nil
}
