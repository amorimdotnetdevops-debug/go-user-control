package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccessControlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateAccessControlHandler",
	})
}

func ShowAccessControlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateAccessControlHandler",
	})
}

func DeleteAccessControlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateAccessControlHandler",
	})
}

func UpdateAccessControlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateAccessControlHandler",
	})
}

func ListAccessControlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateAccessControlHandler",
	})
}
