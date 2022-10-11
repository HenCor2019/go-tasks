package UsersServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) CreateUser(createUserDto dtos.CreateUserDto) models.User {
  savedUser,err := s.repo.CreateUser(createUserDto)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot create the users"))
  }

  return savedUser
}

func (s *Service) FindById(userId string) models.User {
  user,err := s.repo.FindById(userId)
  if err != nil || user.ID == 0 {
    panic(fiber.NewError(fiber.StatusNotFound, "User not found"))
  }

  return user
}

func (s *Service) DeleteById(userIdToDelete string) models.User {
  _,err := s.repo.FindById(userIdToDelete)
  if err != nil {
    panic(fiber.NewError(fiber.StatusNotFound, "User not found"))
  }

  deletedUser,err := s.repo.DeleteById(userIdToDelete)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot delete the user"))
  }

  return deletedUser
}

func (s *Service) Find() []models.User {
  users,err := s.repo.Find()
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot find the users"))
  }

  return users
}
