package repository

import (
	"log"
	"os"
	// "time"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"base/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	log.Println("Running migrations...")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Event{})
}
