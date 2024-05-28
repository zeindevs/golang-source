package controller

import (
	"go-crud-redis/domain"
	"go-crud-redis/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	novelUseCase domain.NovelUseCase
}

func NewNovelController(novelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{novelUseCase: novelUseCase}
}

func (n *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	var novelRequest model.Novel
	var response model.Response

	// handle the request
	if err := ctx.BodyParser(&novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// check if the request was empty/null
	if novelRequest.Author == "" || novelRequest.Name == "" || novelRequest.Description == "" {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "Request cannot be empty"}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// save into database
	if err := n.novelUseCase.CreateNovel(novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

func (n *NovelController) GetNovelById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id (cannot be null / 0)"})
	}

	novel, err := n.novelUseCase.GetNovelById(idInt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var response model.Response
	if novel.Name != "" {
		response = model.Response{StatusCode: http.StatusOK, Message: "Get Novel By Id Succes", Data: novel}
	} else {
		response = model.Response{StatusCode: http.StatusOK, Message: "Get Novel By Id Succes (Null Data)"}
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
