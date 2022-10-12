package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID          uint
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      uint
}
