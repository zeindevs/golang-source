package controller

import (
	"net/http"
	"restful-crud-gin-gorm/data/request"
	"restful-crud-gin-gorm/data/response"
	"restful-crud-gin-gorm/helper"
	"restful-crud-gin-gorm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (c *TagsController) Create(ctx *gin.Context) {
	tagsCreateRequest := request.TagsCreateRequest{}
	err := ctx.ShouldBindJSON(&tagsCreateRequest)
	helper.PanicIfError(err)

	c.tagsService.Create(tagsCreateRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TagsController) Update(ctx *gin.Context) {
	tagsUpdateRequest := request.TagsUpdateRequest{}
	err := ctx.ShouldBindJSON(&tagsUpdateRequest)
	helper.PanicIfError(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)
	tagsUpdateRequest.Id = id

	c.tagsService.Update(tagsUpdateRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)
	c.tagsService.Delete(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)
	tagResponse := c.tagsService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TagsController) FindAll(ctx *gin.Context) {
	tagsResponse := c.tagsService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
