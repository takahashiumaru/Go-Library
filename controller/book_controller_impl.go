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

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(BookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: BookService,
	}
}

func (controller *BookControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "subject.eq")
	bookResponses := controller.BookService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(bookResponses),
		Data:    bookResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {

	request := web.BookCreateRequest{}
	helper.ReadFromRequestBody(c, &request)
	bookResponse := controller.BookService.Create(auth, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Created Successfully",
		Data:    bookResponse,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func (controller *BookControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	book := c.Param("id")
	bookID, err := strconv.Atoi(book)
	helper.PanicIfError(err)
	request := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)
	bookResponse := controller.BookService.Update(auth, &bookID, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Updated Successfully",
		Data:    bookResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	book := c.Param("id")
	bookID, err := strconv.Atoi(book)
	helper.PanicIfError(err)

	controller.BookService.Delete(auth, &bookID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Book deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
