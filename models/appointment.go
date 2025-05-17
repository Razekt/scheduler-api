package models

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Phone string    `json:"phone"`
	Email string    `json:"email"`
}
