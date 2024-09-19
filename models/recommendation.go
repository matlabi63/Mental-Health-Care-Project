package models

import "time"

type Recommendation struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	DoctorID    uint      `json:"doctor_id"`
	UserID      uint      `json:"user_id"`      
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
