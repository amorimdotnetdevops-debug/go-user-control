package router

import (
	"github.com/amorimdotnetdevops-debug/go-user-control/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//initialize handlers
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.ShowUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.DELETE("/user", handler.DeleteUserHandler)
		v1.PUT("/user", handler.UpdateUserHandler)
		v1.GET("/users", handler.ListUserHandler)
	}
}
