package repository

import (
	"go-library/exception"
	"go-library/helper"
	"go-library/model/domain"
	"strings"

	"gorm.io/gorm"
)

type PublisherRepositoryImpl struct {
}

func NewPublisherRepository() PublisherRepository {
	return &PublisherRepositoryImpl{}
}

func (repository *PublisherRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Publishers {
	publishers := domain.Publishers{}
	tx := db.Model(&domain.Publisher{}).
		Joins("CreatedBy").
		Joins("UpdatedBy")
	err := tx.Find(&publishers).Error
	helper.PanicIfError(err)

	return publishers
}

func (repository *PublisherRepositoryImpl) Create(db *gorm.DB, publisher *domain.Publisher) *domain.Publisher {

	err := db.Create(&publisher).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = &exception.ErrorSendToResponse{Err: "record already exists"}
		}
	}
	helper.PanicIfError(err)
	return publisher
}

func (repository *PublisherRepositoryImpl) Update(db *gorm.DB, publisher *domain.Publisher) *domain.Publisher {
	err := db.Updates(&publisher).Error
	helper.PanicIfError(err)

	err = db.First(&publisher).Error
	helper.PanicIfError(err)

	// err = helper.CreateHistory(db, eventDetail, helper.HistoryUpdate, eventDetail.UpdatedByID)
	// helper.PanicIfError(err)

	return publisher
}

func (repository *PublisherRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	eventDetail := &domain.Publisher{}
	tx := db.First(eventDetail, id).Updates(&domain.Publisher{
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
