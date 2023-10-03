package domain

import (
	"time"

	"go-library/model/web"

	"gorm.io/gorm"
)

type Books []Book
type Book struct {
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
	BookCode        string    `gorm:"size:20;not null"`
	Title           string    `gorm:""`
	CategoryID      uint      `gorm:""`
	Category        Category  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET NULL;"`
	PublisherID     uint      `gorm:""`
	Publisher       Publisher `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET NULL;"`
	Isbn            string    `gorm:"size:20"`
	Author          string    `gorm:"size:20"`
	NumberPage      uint      `gorm:""`
	Stock           uint      `gorm:""`
	PublicationYear string    `gorm:"6"`
	Synopsis        string    `gorm:""`
	Image           string    `gorm:""`
}

func (book *Book) ToBookResponse() web.BookResponse {

	return web.BookResponse{
		// Required Fields
		ID:          book.ID,
		CreatedByID: book.CreatedByID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		CreatedBy:   book.CreatedBy.ToUserShortResponse(),
		UpdatedBy:   book.UpdatedBy.ToUserShortResponse(),

		// Fields
		BookCode:        book.BookCode,
		Title:           book.Title,
		CategoryID:      book.CreatedByID,
		PublisherID:     book.PublisherID,
		Isbn:            book.Isbn,
		Author:          book.Author,
		NumberPage:      book.NumberPage,
		Stock:           book.Stock,
		PublicationYear: book.PublicationYear,
		Synopsis:        book.Synopsis,
		Image:           book.Image,
	}
}

func (books Books) ToBookResponses() []web.BookResponse {
	bookResponses := []web.BookResponse{}
	for _, book := range books {
		bookResponses = append(bookResponses, book.ToBookResponse())
	}
	return bookResponses
}
