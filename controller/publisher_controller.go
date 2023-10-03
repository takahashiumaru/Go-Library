package controller

import (
	"go-library/auth"

	"github.com/gin-gonic/gin"
)

type PublisherController interface {
	FindAll(c *gin.Context, auth *auth.AccessDetails)
	Create(c *gin.Context, auth *auth.AccessDetails)
	Update(c *gin.Context, auth *auth.AccessDetails)
	Delete(c *gin.Context, auth *auth.AccessDetails)
}
