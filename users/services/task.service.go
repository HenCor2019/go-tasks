package UsersServices

import (
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/users/dtos"
	"github.com/HenCor2019/task-go/users/repositories"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(createUserDto dtos.CreateUserDto) models.User {
  savedUser,err := UsersRepositories.CreateUser(createUserDto)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot create the users"))
  }

  return savedUser
}

func Find() []models.User {
  users,err := UsersRepositories.Find()
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot find the users"))
  }

  return users
}

func FindById(userId string) models.User {
  user,err := UsersRepositories.FindById(userId)
  if err != nil || user.ID == 0 {
    panic(fiber.NewError(fiber.StatusNotFound, "User not found"))
  }

  return user
}

func DeleteById(userIdToDelete string) models.User {
  _,err := UsersRepositories.FindById(userIdToDelete)
  if err != nil {
    panic(fiber.NewError(fiber.StatusNotFound, "User not found"))
  }

  deletedUser,err := UsersRepositories.DeleteById(userIdToDelete)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot delete the user"))
  }

  return deletedUser
}
