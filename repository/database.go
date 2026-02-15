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


	// dsn := "host=" + os.Getenv("DB_HOST") +
	// 	" user=" + os.Getenv("DB_USER") +
	// 	" password=" + os.Getenv("DB_PASSWORD") +
	// 	" dbname=" + os.Getenv("DB_NAME") +
	// 	" port=" + os.Getenv("DB_PORT") +
	// 	" sslmode=disable"

	// var database *gorm.DB
	// 

	// for i := 0; i < 10; i++ {
	// 	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// 	if err == nil {
	// 		log.Println("Database connected")
	// 		break
	// 	}
	// 	log.Println("Waiting for database...")
	// 	time.Sleep(2 * time.Second)
	// }

	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// database.AutoMigrate(&models.Event{})
	// DB = database
}
