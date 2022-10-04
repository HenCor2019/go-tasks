package db

import (
	"log"

	"github.com/HenCor2019/task-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION_STRING = "host=localhost user=hencor password=password dbname=taskgodb port=5432"
var DB *gorm.DB

func DBConnection(){

  var err error
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

  log.Println("Connected to the database")
}
