package repository

import (
	"go-library/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Categorys
	Create(db *gorm.DB, category *domain.Category) *domain.Category
	Update(db *gorm.DB, category *domain.Category) *domain.Category
	Delete(db *gorm.DB, categoryID *int, deletedByID *uint)
}
