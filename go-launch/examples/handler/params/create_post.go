package params

// Request domain
type CreatePost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p CreatePost) Validate() (any, bool) {
	if len(p.Title) < 100 {
		return map[string]string{"title": "to short"}, false
	}
	return nil, true
}
