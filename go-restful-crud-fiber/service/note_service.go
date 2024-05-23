package service

import (
	"restful-crud-fiber/data/request"
	"restful-crud-fiber/data/response"
)

type NoteService interface {
	Create(note request.NoteCreateRequest)
	Update(note request.NoteUpdateRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
