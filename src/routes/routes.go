package routes

import (
	"github.com/gin-gonic/gin"

	content_rest "go-api/content/rest"
)

func Run(address string) error {
	router := gin.Default()
	v1 := router.Group("/")
	content_rest.AddContentRoutes(v1)
	return router.Run(address)
}
