package handler

type CreateUserRequest struct {
	Role string `json:"role" binding:"required"`
}