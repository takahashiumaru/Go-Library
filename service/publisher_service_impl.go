package service

import (
	"go-library/auth"
	"go-library/helper"
	"go-library/model/domain"
	"go-library/model/web"
	"go-library/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PublisherServiceImpl struct {
	PublisherRepository repository.PublisherRepository
	DB                  *gorm.DB
	Validate            *validator.Validate
}

func NewPublisherService(
	publisher repository.PublisherRepository,
	db *gorm.DB,
	validate *validator.Validate,
) PublisherService {
	return &PublisherServiceImpl{
		PublisherRepository: publisher,
		DB:                  db,
		Validate:            validate,
	}
}

func (service *PublisherServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.PublisherResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	publishers := service.PublisherRepository.FindAll(tx, filters)
	return publishers.ToPublisherResponses()
}

func (service *PublisherServiceImpl) Create(auth *auth.AccessDetails, request *web.PublisherCreateRequest) web.PublisherResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	publisher := &domain.Publisher{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		PublisherCode: request.PublisherCode,
		Name:          request.Name,
	}
	publisherResponse := service.PublisherRepository.Create(tx, publisher)
	return publisherResponse.ToPublisherResponse()
}

func (service *PublisherServiceImpl) Update(auth *auth.AccessDetails, publisherID *int, request *web.PublisherUpdateRequest) web.PublisherResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	publisher := &domain.Publisher{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		ID:            uint(*publisherID),
		PublisherCode: request.PublisherCode,
		Name:          request.Name,
	}
	publisherResponse := service.PublisherRepository.Update(tx, publisher)
	return publisherResponse.ToPublisherResponse()
}

func (service *PublisherServiceImpl) Delete(auth *auth.AccessDetails, publisherID *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	service.PublisherRepository.Delete(tx, publisherID, &auth.UserID)
	defer helper.CommitOrRollback(tx)
}
