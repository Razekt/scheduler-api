package models

type AppointmentResponse struct {
	Code    int
	Message string
}

func NewAppointmentResponse(code int, message string) AppointmentResponse {
	return AppointmentResponse{
		Code:    code,
		Message: message,
	}
}
