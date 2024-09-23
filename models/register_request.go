package models

import (
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Role     string `json:"role" validate:"required,min=2,max=50"`
}

func (r *RegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
