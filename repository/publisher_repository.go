package repository

import (
	"go-library/model/domain"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Publishers
	Create(db *gorm.DB, publisher *domain.Publisher) *domain.Publisher
	Update(db *gorm.DB, publisher *domain.Publisher) *domain.Publisher
	Delete(db *gorm.DB, publisherID *int, deletedByID *uint)
}
