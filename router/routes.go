package router

import (
	"github.com/amorimdotnetdevops-debug/go-user-control/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handler.ShowAccessControlHandler)
		v1.POST("/ping", handler.CreateUserHandler)
		v1.DELETE("/ping", handler.DeleteAccessControlHandler)
		v1.PUT("/ping", handler.UpdateAccessControlHandler)
		v1.GET("/pings", handler.ListAccessControlHandler)
	}
}
