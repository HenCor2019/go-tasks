package db

import (
	"fmt"
	"log"
	"os"

	"github.com/HenCor2019/task-go/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dbConfig := os.Getenv

	CONNECTION_STRING := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig("DB_HOST"),
		dbConfig("DB_USERNAME"),
		dbConfig("DB_PASSWORD"),
		dbConfig("DB_NAME"),
		dbConfig("DB_PORT"),
	)
	DB, err := gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to the database")
		panic(err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Error connection to the database")
		panic(err)
	}

	log.Println("Connected to the database successfully")
	return DB, nil
}
