package models

import (
	"github.com/go-playground/validator/v10"
)

type AdminRequest struct {
	UserID      uint  `json:"user_id" validate:"required"`
	ManageUsers bool  `json:"manage_users" validate:"required"`
}

func (a *AdminRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
