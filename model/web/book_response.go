package web

import (
	"time"
)

type BookResponse struct {
	// Required Fields
	ID          uint              `json:"id"`
	CreatedByID uint              `json:"created_by_id"`
	UpdatedByID uint              `json:"updated_by_id"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	CreatedBy   UserShortResponse `json:"created_by"`
	UpdatedBy   UserShortResponse `json:"updated_by"`

	// Required Fields
	BookCode        string `json:"book_code"`
	Title           string `json:"title"`
	CategoryID      uint   `json:"category_id"`
	PublisherID     uint   `json:"publisher_id"`
	Isbn            string `json:"isbn"`
	Author          string `json:"author"`
	NumberPage      uint   `json:"number_page"`
	Stock           uint   `json:"stock"`
	PublicationYear string `json:"publication_year"`
	Synopsis        string `json:"synopsis"`
	Image           string `json:"image"`
}
