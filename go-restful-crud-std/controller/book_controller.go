package controller

import (
	"net/http"
	"restful-crud-std/data/request"
	"restful-crud-std/data/response"
	"restful-crud-std/helper"
	"restful-crud-std/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(r, &bookCreateRequest)

	c.BookService.Create(r.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(r, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	c.BookService.Update(r.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	c.BookService.Delete(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (c *BookController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	book := c.BookService.FindById(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   book,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	books := c.BookService.FindAll(r.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   books,
	}

	helper.WriteResponseBody(w, webResponse)

}
