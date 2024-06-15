package service

type TagsService interface {
	Create(tags any)
	Update(tags any)
	Delete(tagsId int)
	FindById(tagsId int) any
	FindAll() []any
}
