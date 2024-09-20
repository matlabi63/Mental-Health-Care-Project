package models

import "time"

type Doctor struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Specialty string    `json:"specialty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}