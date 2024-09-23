package models

import (
	"github.com/go-playground/validator/v10"
)

type RecommendationRequest struct {
	Description string `json:"description" validate:"required,min=10,max=500"`
	DoctorID    uint   `json:"doctor_id" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}

func (r *RecommendationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
