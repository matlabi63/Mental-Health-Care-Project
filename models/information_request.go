package models

import (
	"github.com/go-playground/validator/v10"
)

type InformationRequest struct {
	Title   string `json:"title" validate:"required,min=5,max=255"`
	Content string `json:"content" validate:"required,min=10"`
	Author  string `json:"author" validate:"required,min=2,max=255"`
}

func (i *InformationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}
