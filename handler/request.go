package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateUserRequest struct {
	// Role string `json:"role" binding:"required"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	EMail  string `json:"email"`
	Active *bool  `json:"active"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.EMail == "" && r.Active == nil {
		return fmt.Errorf("request body is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.EMail == "" {
		return errParamIsRequired("EMail", "string")
	}
	if r.Active == nil {
		return errParamIsRequired("Active", "bool")
	}
	return nil
}
