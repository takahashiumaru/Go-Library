package route

import (
	"go-library/auth"
	"go-library/controller"
	"go-library/repository"
	"go-library/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CategoryRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	categoryService := service.NewCategoryService(
		repository.NewCategoryRepository(),
		db,
		validate,
	)

	categoryRoute := controller.NewCategoryController(categoryService)
	router.GET("/category", auth.Auth(categoryRoute.FindAll, []string{}))
	router.POST("/category", auth.Auth(categoryRoute.Create, []string{}))
	router.PUT("/category/:id", auth.Auth(categoryRoute.Update, []string{}))
	router.DELETE("/category/:id", auth.Auth(categoryRoute.Delete, []string{}))
}
