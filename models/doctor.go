package models

import "time"

type Doctor struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Name      string    `json:"name"`
	Specialty string    `json:"specialty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
