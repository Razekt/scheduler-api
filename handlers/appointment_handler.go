package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/scheduler-api/models"
	"github.com/scheduler-api/validators"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppointmentHandler struct {
	DB *gorm.DB
}

func NewAppointmentHandler(db *gorm.DB) *AppointmentHandler {
	return &AppointmentHandler{DB: db}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var input models.AppointmentInput
	var response models.AppointmentResponse = models.NewAppointmentResponse(1, "Appointment created")
	if err := c.ShouldBindJSON(&input); err != nil {
		response = models.NewAppointmentResponse(0, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if valid, msg := validators.ValidateAppointmentInput(h.DB, input); !valid {
		response = models.NewAppointmentResponse(-6, msg)
		c.JSON(http.StatusConflict, response)
		return
	}

	a := models.Appointment{
		ID:    uuid.New(),
		Name:  input.Name,
		Date:  input.Date,
		Phone: input.Phone,
		Email: input.Email,
	}

	if err := h.DB.Create(&a).Error; err != nil {
		response = models.NewAppointmentResponse(-1, "Could not create an appointment")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AppointmentHandler) ListAppointments(c *gin.Context) {
	var list []models.Appointment
	if err := h.DB.Limit(10).Find(&list).Error; err != nil {
		response := models.NewAppointmentResponse(0, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id := c.Param("id")
	var response models.AppointmentResponse
	var a models.Appointment

	if _, err := uuid.Parse(id); err != nil {
		response = models.NewAppointmentResponse(-3, "Invalid UUID format")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.DB.First(&a, "id = ?", id).Error; err != nil {
		response = models.NewAppointmentResponse(-2, "Appointment not found")
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var response models.AppointmentResponse = models.NewAppointmentResponse(2, "Appointment deleted!")
	if err := h.DB.Delete(&models.Appointment{}, "id = ?", id).Error; err != nil {
		response = models.NewAppointmentResponse(-4, "Could not delete appointment!")
	}
	c.JSON(http.StatusOK, response)
}
