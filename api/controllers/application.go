package controllers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@BasePath	/api/v1/

// CreateApplication godoc
//	@Summary	Create an application
//	@Schemes
//	@Description	create an application
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/apps/ [post]
func CreateApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(501, "Not implemented yet")
	}
}

// GetApplication godoc
//	@Summary	Get all applications created by the user
//	@Schemes
//	@Description	Get all applications created by the user
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/apps/ [get]
func GetApplications() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(501, "Not implemented yet")
	}
}