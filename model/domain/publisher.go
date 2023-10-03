package domain

import (
	"time"

	"go-library/model/web"

	"gorm.io/gorm"
)

type Publishers []Publisher
type Publisher struct {
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
	PublisherCode string `gorm:"size:20;not null"`
	Name          string `gorm:""`
}

func (publisher *Publisher) ToPublisherResponse() web.PublisherResponse {

	return web.PublisherResponse{
		// Required Fields
		ID:          publisher.ID,
		CreatedByID: publisher.CreatedByID,
		CreatedAt:   publisher.CreatedAt,
		UpdatedAt:   publisher.UpdatedAt,
		CreatedBy:   publisher.CreatedBy.ToUserShortResponse(),
		UpdatedBy:   publisher.UpdatedBy.ToUserShortResponse(),

		// Fields
		PublisherCode: publisher.PublisherCode,
		Name:          publisher.Name,
	}
}

func (publishers Publishers) ToPublisherResponses() []web.PublisherResponse {
	PublisherResponses := []web.PublisherResponse{}
	for _, publisher := range publishers {
		PublisherResponses = append(PublisherResponses, publisher.ToPublisherResponse())
	}
	return PublisherResponses
}
