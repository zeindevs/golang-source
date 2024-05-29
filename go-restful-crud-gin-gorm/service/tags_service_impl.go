package service

import (
	"restful-crud-gin-gorm/data/request"
	"restful-crud-gin-gorm/data/response"
	"restful-crud-gin-gorm/helper"
	"restful-crud-gin-gorm/model"
	"restful-crud-gin-gorm/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagsRepository,
		Validate:       validate,
	}
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.TagsCreateRequest) {
	err := t.Validate.Struct(tags)
	helper.PanicIfError(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}

		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.PanicIfError(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse

}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.TagsUpdateRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.PanicIfError(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
