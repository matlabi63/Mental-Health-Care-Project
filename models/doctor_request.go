package models

import "github.com/go-playground/validator/v10"

type DoctorRequest struct {
	UserID    uint   `json:"user_id" validate:"required"`
	Name      string `json:"name" validate:"required,min=2,max=50"`
	Specialty string `json:"specialty" validate:"required,min=2,max=50"`
}

func (d *DoctorRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
