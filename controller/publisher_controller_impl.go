package controller

import (
	"net/http"
	"strconv"

	"go-library/auth"
	"go-library/helper"
	"go-library/model/web"
	"go-library/service"

	"github.com/gin-gonic/gin"
)

type PublisherControllerImpl struct {
	PublisherService service.PublisherService
}

func NewPublisherController(PublisherService service.PublisherService) PublisherController {
	return &PublisherControllerImpl{
		PublisherService: PublisherService,
	}
}

func (controller *PublisherControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "subject.eq")
	publisherResponses := controller.PublisherService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(publisherResponses),
		Data:    publisherResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *PublisherControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {

	request := web.PublisherCreateRequest{}
	helper.ReadFromRequestBody(c, &request)
	publisherResponse := controller.PublisherService.Create(auth, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "publisher Publisher Created Successfully",
		Data:    publisherResponse,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func (controller *PublisherControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	publisher := c.Param("id")
	publisherID, err := strconv.Atoi(publisher)
	helper.PanicIfError(err)
	request := web.PublisherUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)
	publisherResponse := controller.PublisherService.Update(auth, &publisherID, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Publisher Updated Successfully",
		Data:    publisherResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *PublisherControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	publisher := c.Param("id")
	publisherID, err := strconv.Atoi(publisher)
	helper.PanicIfError(err)

	controller.PublisherService.Delete(auth, &publisherID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Publisher deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
