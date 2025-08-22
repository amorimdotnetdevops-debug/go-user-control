package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// initialize the Gin router
	router := gin.Default()

	//Initialize the routes
	initializeRoutes(router)

	// run the server
	// This will block until the server is stopped
	router.Run(":8080") // listen and serve on :8080
}
