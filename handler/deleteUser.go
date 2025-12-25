package handler

import (
	"fmt"
	"net/http"

	"github.com/amorimdotnetdevops-debug/go-user-control/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("Failed to found user: %v", err.Error())
		sendError(ctx, http.StatusFound, fmt.Sprintf("User with id %s not found", id))
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		message := fmt.Sprintf("Failed to delete user with id %s. Error: %v", id, err.Error())
		logger.Errorf("%v", message)

		sendError(ctx, http.StatusInternalServerError, message)
		return
	}

	sendSuccess(ctx, "Delete-User", user)
}
