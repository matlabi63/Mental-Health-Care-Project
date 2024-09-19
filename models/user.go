package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255" json:"name"`
	Email     string         `gorm:"size:255;unique" json:"email"`
	Password  string         `gorm:"size:255" json:"password"`
	Role      string         `gorm:"size:50" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
