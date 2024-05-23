package request

type NoteCreateRequest struct {
	Content string `validate:"required,min=2,max=100" json:"content"`
}

type NoteUpdateRequest struct {
	Id      int    `validate:"required"`
	Content string `validate:"required,min=2,max=100" json:"content"`
}
