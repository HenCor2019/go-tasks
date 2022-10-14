package TasksServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/tasks/dtos"
	"github.com/gofiber/fiber/v2"
)

func (taskServices *Service) CreateTask(createTaskDto tasksDtos.CreateTaskDto, userId string) models.Task {
	user, err := taskServices.repo.FindUserById(userId)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, "User not found"))
	}
	savedTask, err := taskServices.repo.Create(createTaskDto, user.ID)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, "Cannot save the task"))
	}

	return savedTask
}

func (taskServices *Service) DeleteTaskById(taskId string, userId string) {
	user, err := taskServices.repo.FindUserById(userId)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, "User not found"))
	}

	_, err = taskServices.repo.FindById(taskId, user.ID)
	if err != nil {
		panic(fiber.NewError(fiber.StatusForbidden, "You cannot delete this task"))
	}

	err = taskServices.repo.DeleteById(taskId)

	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, "Cannot Delete the task"))
	}
}
