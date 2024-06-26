package service

import (
	"context"
	"restful-crud-std/data/request"
	"restful-crud-std/data/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
