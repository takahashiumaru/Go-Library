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

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(CategoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: CategoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "subject.eq")
	categoryResponses := controller.CategoryService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(categoryResponses),
		Data:    categoryResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {

	request := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(c, &request)
	categoryResponse := controller.CategoryService.Create(auth, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "category Category Created Successfully",
		Data:    categoryResponse,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func (controller *CategoryControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	category := c.Param("id")
	categoryID, err := strconv.Atoi(category)
	helper.PanicIfError(err)
	request := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)
	categoryResponse := controller.CategoryService.Update(auth, &categoryID, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Category Updated Successfully",
		Data:    categoryResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	category := c.Param("id")
	categoryID, err := strconv.Atoi(category)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(auth, &categoryID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Book Category deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
