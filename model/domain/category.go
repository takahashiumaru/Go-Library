package domain

import (
	"time"

	"go-library/model/web"

	"gorm.io/gorm"
)

type Categorys []Category
type Category struct {
	// Required Fields
	gorm.Model
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:""`
	CreatedByID uint      `gorm:""`
	UpdatedAt   time.Time `gorm:""`
	UpdatedByID uint      `gorm:""`
	DeletedByID *uint     `gorm:""`
	CreatedBy   User
	UpdatedBy   User

	// Fields
	CategoryCode string `gorm:"size:20;not null"`
	Name         string `gorm:""`
}

func (category *Category) ToCategoryResponse() web.CategoryResponse {

	return web.CategoryResponse{
		// Required Fields
		ID:          category.ID,
		CreatedByID: category.CreatedByID,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
		CreatedBy:   category.CreatedBy.ToUserShortResponse(),
		UpdatedBy:   category.UpdatedBy.ToUserShortResponse(),

		// Fields
		CategoryCode: category.CategoryCode,
		Name:         category.Name,
	}
}

func (categorys Categorys) ToCategoryResponses() []web.CategoryResponse {
	CategoryResponses := []web.CategoryResponse{}
	for _, category := range categorys {
		CategoryResponses = append(CategoryResponses, category.ToCategoryResponse())
	}
	return CategoryResponses
}
