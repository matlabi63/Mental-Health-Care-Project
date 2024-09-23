package models

import (
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}
