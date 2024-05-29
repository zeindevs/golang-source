package request

type TagsUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}
