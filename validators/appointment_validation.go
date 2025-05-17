package validators

import (
	"github.com/scheduler-api/models"
	"gorm.io/gorm"
)

func ValidateAppointmentInput(db *gorm.DB, input models.AppointmentInput) (bool, string) {
	var existing models.Appointment
	err := db.Where("email = ? AND date = ?", input.Email, input.Date).First(&existing).Error
	if err == nil && existing.Email != "" && !existing.Date.IsZero() {
		return false, "Você já possui um agendamento nesse horário"
	}

	err = db.Where("date = ?", input.Date).First(&existing).Error
	if err == nil && !existing.Date.IsZero() {
		return false, "Este horário já está reservado"
	}

	return true, ""
}
