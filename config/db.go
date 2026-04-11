package config

import (
	"fmt"
	"log"
	"noteapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=akzharkyn dbname=noteapp port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Category{}, &models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = database
	fmt.Println("Database connected successfully")
}
