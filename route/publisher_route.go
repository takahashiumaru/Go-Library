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

func PublisherRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	publisherService := service.NewPublisherService(
		repository.NewPublisherRepository(),
		db,
		validate,
	)

	publisherRoute := controller.NewPublisherController(publisherService)
	router.GET("/publisher", auth.Auth(publisherRoute.FindAll, []string{}))
	router.POST("/publisher", auth.Auth(publisherRoute.Create, []string{}))
	router.PUT("/publisher/:id", auth.Auth(publisherRoute.Update, []string{}))
	router.DELETE("/publisher/:id", auth.Auth(publisherRoute.Delete, []string{}))
}
