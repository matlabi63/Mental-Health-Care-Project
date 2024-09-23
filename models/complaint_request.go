package models

import (
	"time"
)

type ComplaintRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Details string `json:"details" validate:"required,min=10,max=500"`
}

func (cr ComplaintRequest) ToComplaint() Complaint {
	return Complaint{
		UserID:    cr.UserID,
		Details:   cr.Details,
		CreatedAt: time.Now(), // You can also set the other fields as needed
		UpdatedAt: time.Now(),
	}
}

// func (c *ComplaintRequest) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(c)
// }


