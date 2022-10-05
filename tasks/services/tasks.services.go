package TasksServices

import (
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/tasks/dtos"
	"github.com/HenCor2019/task-go/tasks/repositories"
	"github.com/HenCor2019/task-go/users/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(createTaskDto tasksDtos.CreateTaskDto, userId string) models.Task {
  user := UsersServices.FindById(userId)
  savedTask,err := TasksRepositories.Create(createTaskDto,user.ID)

  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest,"Cannot save the task"))
  }

  return savedTask
}

func DeleteTask(taskId string, userId string) {
  user := UsersServices.FindById(userId)

  _,err := TasksRepositories.FindById(taskId,user.ID)
  if err != nil {
    panic(fiber.NewError(fiber.StatusForbidden,"You cannot delete this task"))
  }

  err = TasksRepositories.DeleteById(taskId)

  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest,"Cannot Delete the task"))
  }
}
