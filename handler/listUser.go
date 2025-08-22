package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ListUserHandler",
	})
}
