package service

import (
	"go-library/auth"
	"go-library/model/web"
)

type PublisherService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.PublisherResponse
	Create(auth *auth.AccessDetails, request *web.PublisherCreateRequest) web.PublisherResponse
	Update(auth *auth.AccessDetails, publisherID *int, request *web.PublisherUpdateRequest) web.PublisherResponse
	Delete(auth *auth.AccessDetails, publisherID *int)
}
