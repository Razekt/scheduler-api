package main

import (
	"github.com/scheduler-api/database"
	"github.com/scheduler-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	r := gin.Default()
	handler := handlers.NewAppointmentHandler(db)

	r.POST("/appointments", handler.CreateAppointment)
	r.GET("/appointments/all", handler.ListAppointments)
	r.GET("/appointments/:id", handler.GetAppointment)
	r.DELETE("/appointments/:id", handler.DeleteAppointment)

	r.Run(":8080")
}
