package service

import (
	"go-library/auth"
	"go-library/model/web"
)

type BookService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.BookResponse
	Create(auth *auth.AccessDetails, request *web.BookCreateRequest) web.BookResponse
	Update(auth *auth.AccessDetails, bookID *int, request *web.BookUpdateRequest) web.BookResponse
	Delete(auth *auth.AccessDetails, bookID *int)
}
