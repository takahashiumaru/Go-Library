package repository

import (
	"go-library/model/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Books
	Create(db *gorm.DB, book *domain.Book) *domain.Book
	Update(db *gorm.DB, book *domain.Book) *domain.Book
	Delete(db *gorm.DB, bookID *int, deletedByID *uint)
}
