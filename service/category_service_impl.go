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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCategoryService(
	category repository.CategoryRepository,
	db *gorm.DB,
	validate *validator.Validate,
) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: category,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.CategoryResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	books := service.CategoryRepository.FindAll(tx, filters)
	return books.ToCategoryResponses()
}

func (service *CategoryServiceImpl) Create(auth *auth.AccessDetails, request *web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	book := &domain.Category{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		CategoryCode: request.CategoryCode,
		Name:         request.Name,
	}
	CategoryResponse := service.CategoryRepository.Create(tx, book)
	return CategoryResponse.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Update(auth *auth.AccessDetails, bookID *int, request *web.CategoryUpdateRequest) web.CategoryResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	book := &domain.Category{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		ID:           uint(*bookID),
		CategoryCode: request.CategoryCode,
		Name:         request.Name,
	}
	CategoryResponse := service.CategoryRepository.Update(tx, book)
	return CategoryResponse.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Delete(auth *auth.AccessDetails, bookID *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	service.CategoryRepository.Delete(tx, bookID, &auth.UserID)
	defer helper.CommitOrRollback(tx)
}
