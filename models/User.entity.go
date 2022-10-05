package models

import "gorm.io/gorm"

type User struct {
  gorm.Model

  Name string `gorm:"not null" json:"name"`
  Email string`gorm:"not null;unique_index" json:"email"`
  Age uint8 `gorm:"not null" json:"age"`
  Tasks []Task `json:"tasks"`
}
