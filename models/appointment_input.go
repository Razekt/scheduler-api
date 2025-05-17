package models

import "time"

type AppointmentInput struct {
	Name  string    `json:"name" binding:"required"`
	Date  time.Time `json:"date" binding:"required"`
	Phone string    `json:"phone" binding:"required"`
	Email string    `json:"email" binding:"email"`
}
