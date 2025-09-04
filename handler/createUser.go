package handler

import (
	"net/http"

	"github.com/amorimdotnetdevops-debug/go-user-control/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %+v", request)

	// if err := db.Exec("INSERT INTO users (role) VALUES (?)", request.Role).Error; err != nil {
	// 	logger.Errorf("Failed to create user: %v", err)
	// 	ctx.JSON(500, gin.H{"error": "Failed to create user"})
	// 	return
	// }

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name:   request.Name,
		Email:  request.EMail,
		Active: request.Active,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("Failed to create user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}

	sendSuccess(ctx, "Create-User", user)
}
