package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := struct {
		Role string `json:"role"`
	}{}

	ctx.BindJSON(&request)

	// if err := db.Exec("INSERT INTO users (role) VALUES (?)", request.Role).Error; err != nil {
	// 	logger.Errorf("Failed to create user: %v", err)
	// 	ctx.JSON(500, gin.H{"error": "Failed to create user"})
	// 	return
	// }

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Failed to create user: %v", err)
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
}
