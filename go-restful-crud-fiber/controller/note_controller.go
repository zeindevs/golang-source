package controller

import (
	"restful-crud-fiber/data/request"
	"restful-crud-fiber/data/response"
	"restful-crud-fiber/helper"
	"restful-crud-fiber/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (c *NoteController) Create(ctx *fiber.Ctx) error {
	noteCreateRequest := request.NoteCreateRequest{}
	err := ctx.BodyParser(&noteCreateRequest)
	helper.ErrorPanic(err)

	c.noteService.Create(noteCreateRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c *NoteController) Update(ctx *fiber.Ctx) error {
	noteUpdateRequest := request.NoteUpdateRequest{}
	err := ctx.BodyParser(&noteUpdateRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	noteUpdateRequest.Id = id

	c.noteService.Update(noteUpdateRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	c.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	noteResponse := c.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get one notes data!",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c *NoteController) FindAll(ctx *fiber.Ctx) error {
	notesResponse := c.noteService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get all notes data!",
		Data:    notesResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
