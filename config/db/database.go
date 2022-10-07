package db

import (
	"fmt"
	"log"
	"os"

	"github.com/HenCor2019/task-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection(){

  var err error
  DB_NAME := os.Getenv("DB_NAME")
  DB_HOST := os.Getenv("DB_HOST")
  DB_USERNAME := os.Getenv("DB_USERNAME")
  DB_PASSWORD := os.Getenv("DB_PASSWORD")
  DB_PORT := os.Getenv("DB_PORT")
  CONNECTION_STRING := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST,DB_USERNAME,DB_PASSWORD,DB_NAME,DB_PORT)
  DB, err = gorm.Open(postgres.Open(CONNECTION_STRING ), &gorm.Config{})

  if err != nil {
    log.Fatal("Cannot connect to the database")
    panic(err)
  }

  err = DB.AutoMigrate(&models.User {}, &models.Task {})
  if err != nil {
    log.Fatal("Error connection to the database")
    panic(err)
  }

  log.Println("Connected to the database successfully")
}
