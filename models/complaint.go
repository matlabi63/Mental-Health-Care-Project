package models

import "time"

type Complaint struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
