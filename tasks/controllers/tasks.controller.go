package TasksController

import (
	"github.com/HenCor2019/task-go/common/responses"
	"github.com/HenCor2019/task-go/tasks/dtos"
	"github.com/HenCor2019/task-go/tasks/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
  createTaskDto := tasksDtos.CreateTaskDto {}
  c.BodyParser(&createTaskDto)

  userId := c.Params("id")
  savedTask := TasksServices.CreateTask(createTaskDto,userId)

  return common.SuccessResponse(c,savedTask,fiber.StatusCreated)
}

func DeleteTask(c *fiber.Ctx) error {
  userId := c.Params("userId")
  taskId := c.Params("taskId")

  TasksServices.DeleteTask(taskId,userId)
  return common.SuccessResponse(c,nil,fiber.StatusNoContent)
}
