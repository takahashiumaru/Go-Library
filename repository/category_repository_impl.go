package repository

import (
	"go-library/exception"
	"go-library/helper"
	"go-library/model/domain"
	"strings"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Categorys {
	categorys := domain.Categorys{}
	tx := db.Model(&domain.Category{}).
		Joins("CreatedBy").
		Joins("UpdatedBy")
	err := tx.Find(&categorys).Error
	helper.PanicIfError(err)

	return categorys
}

func (repository *CategoryRepositoryImpl) Create(db *gorm.DB, category *domain.Category) *domain.Category {

	err := db.Create(&category).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = &exception.ErrorSendToResponse{Err: "record already exists"}
		}
	}
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(db *gorm.DB, category *domain.Category) *domain.Category {
	err := db.Updates(&category).Error
	helper.PanicIfError(err)

	err = db.First(&category).Error
	helper.PanicIfError(err)

	// err = helper.CreateHistory(db, eventDetail, helper.HistoryUpdate, eventDetail.UpdatedByID)
	// helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	eventDetail := &domain.Category{}
	tx := db.First(eventDetail, id).Updates(&domain.Category{
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
