package service

import (
	"go-library/auth"
	"go-library/model/web"
)

type CategoryService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.CategoryResponse
	Create(auth *auth.AccessDetails, request *web.CategoryCreateRequest) web.CategoryResponse
	Update(auth *auth.AccessDetails, categoryID *int, request *web.CategoryUpdateRequest) web.CategoryResponse
	Delete(auth *auth.AccessDetails, categoryID *int)
}
