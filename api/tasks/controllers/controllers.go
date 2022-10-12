package TasksControllers

import (
	"github.com/HenCor2019/task-go/api/tasks/services"
	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	CreateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type Controller struct {
	service TasksServices.TaskService
}

func New(service TasksServices.TaskService) TaskController {
	return &Controller{ service }
}
