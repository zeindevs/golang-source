package service

import (
	"restful-crud-gin-gorm/data/request"
	"restful-crud-gin-gorm/data/response"
)

type TagsService interface {
	Create(tags request.TagsCreateRequest)
	Update(tags request.TagsUpdateRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
