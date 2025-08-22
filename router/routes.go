package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET",
			})
		})
		v1.POST("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "POST",
			})
		})
		v1.DELETE("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE",
			})
		})
		v1.PUT("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT",
			})
		})
		v1.GET("/pings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GETs",
			})
		})
	}
}
