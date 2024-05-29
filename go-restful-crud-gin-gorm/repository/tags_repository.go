package repository

import "restful-crud-gin-gorm/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsDd int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
