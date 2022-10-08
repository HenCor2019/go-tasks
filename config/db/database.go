package db

import (
	"fmt"
	"log"

	"github.com/HenCor2019/task-go/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection(){

  var err error

  dbConfig := viper.Get

  CONNECTION_STRING := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
                                    dbConfig("DB_HOST"),
                                    dbConfig("DB_USERNAME"),
                                    dbConfig("DB_PASSWORD"),
                                    dbConfig("DB_NAME"),
                                    dbConfig("DB_PORT"),
                       )
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
