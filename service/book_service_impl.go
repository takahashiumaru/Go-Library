package service

import (
	"go-library/auth"
	"go-library/helper"
	"go-library/model/domain"
	"go-library/model/web"
	"go-library/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewBookService(
	book repository.BookRepository,
	db *gorm.DB,
	validate *validator.Validate,
) BookService {
	return &BookServiceImpl{
		BookRepository: book,
		DB:             db,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.BookResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	books := service.BookRepository.FindAll(tx, filters)
	return books.ToBookResponses()
}

func (service *BookServiceImpl) Create(auth *auth.AccessDetails, request *web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	book := &domain.Book{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		BookCode: request.BookCode,
		Title:    request.Title,
	}
	bookResponse := service.BookRepository.Create(tx, book)
	return bookResponse.ToBookResponse()
}

func (service *BookServiceImpl) Update(auth *auth.AccessDetails, bookID *int, request *web.BookUpdateRequest) web.BookResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	book := &domain.Book{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		ID:       uint(*bookID),
		BookCode: request.BookCode,
		Title:    request.Title,
	}
	bookResponse := service.BookRepository.Update(tx, book)
	return bookResponse.ToBookResponse()
}

func (service *BookServiceImpl) Delete(auth *auth.AccessDetails, bookID *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	service.BookRepository.Delete(tx, bookID, &auth.UserID)
	defer helper.CommitOrRollback(tx)
}
