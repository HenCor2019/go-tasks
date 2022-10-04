package models

type Task struct {
  Title string `gorm:"not null"`
  Description string `gorm:"not null"`
  UserID  uint
}
