package params

import "fmt"

// Request domain
type CreatePost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p CreatePost) Validate() error {
	if len(p.Title) < 100 {
		return fmt.Errorf("invalid email")
	}
	return nil
}
