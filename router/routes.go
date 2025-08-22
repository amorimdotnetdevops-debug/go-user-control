package router

import (
	"github.com/amorimdotnetdevops-debug/go-user-control/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.ShowUserHandler)
		v1.POST("/ping", handler.CreateUserHandler)
		v1.DELETE("/ping", handler.DeleteUserHandler)
		v1.PUT("/ping", handler.UpdateUserHandler)
		v1.GET("/pings", handler.ListUserHandler)
	}
}
