package models

import "github.com/go-playground/validator/v10"

type ComplaintRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Details string `json:"details" validate:"required,min=10,max=500"`
}

func (c *ComplaintRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	return err
}