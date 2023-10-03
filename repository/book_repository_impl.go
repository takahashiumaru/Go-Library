package repository

import (
	"go-library/exception"
	"go-library/helper"
	"go-library/model/domain"
	"strings"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Books {
	books := domain.Books{}
	tx := db.Model(&domain.Book{}).
		Joins("CreatedBy").
		Joins("UpdatedBy")
	err := tx.Find(&books).Error
	helper.PanicIfError(err)

	return books
}

func (repository *BookRepositoryImpl) Create(db *gorm.DB, book *domain.Book) *domain.Book {

	err := db.Create(&book).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = &exception.ErrorSendToResponse{Err: "record already exists"}
		}
	}
	helper.PanicIfError(err)
	return book
}

func (repository *BookRepositoryImpl) Update(db *gorm.DB, book *domain.Book) *domain.Book {
	err := db.Updates(&book).Error
	helper.PanicIfError(err)

	err = db.First(&book).Error
	helper.PanicIfError(err)

	// err = helper.CreateHistory(db, eventDetail, helper.HistoryUpdate, eventDetail.UpdatedByID)
	// helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	eventDetail := &domain.Book{}
	tx := db.First(eventDetail, id).Updates(&domain.Book{
		// Model:       gorm.Model{ID: uint(*id)},
		DeletedByID: deletedByID,
	})

	// Creating a history of the deleted event detail.
	// err := helper.CreateHistory(db, eventDetail, helper.HistoryDelete, *deletedByID)
	// helper.PanicIfError(err)

	// Deleting the event detail from the database.
	err := tx.Unscoped().Delete(eventDetail, id).Error
	helper.PanicIfError(err)
}
