package models

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Role     string `json:"role" validate:"required,min=2,max=50"`
}