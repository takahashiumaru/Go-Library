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

func BookRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	bookService := service.NewBookService(
		repository.NewBookRepository(),
		db,
		validate,
	)

	bookRoute := controller.NewBookController(bookService)
	router.GET("/book", auth.Auth(bookRoute.FindAll, []string{}))
	router.POST("/book", auth.Auth(bookRoute.Create, []string{}))
	router.PUT("/book/:id", auth.Auth(bookRoute.Update, []string{}))
	router.DELETE("/book/:id", auth.Auth(bookRoute.Delete, []string{}))
}
