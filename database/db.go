package database

import (
	"github.com/scheduler-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbManager, err := gorm.Open(sqlite.Open("scheduler.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	dbManager.AutoMigrate(&models.Appointment{})
	return dbManager
}
