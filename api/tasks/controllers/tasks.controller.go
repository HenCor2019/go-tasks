package TasksControllers

import (
	"github.com/HenCor2019/task-go/api/common/responses"
	tasksDtos "github.com/HenCor2019/task-go/api/tasks/dtos"
	"github.com/gofiber/fiber/v2"
)

func (taskController *Controller) CreateTask(c *fiber.Ctx) error {
	createTaskDto := tasksDtos.CreateTaskDto{}
	err := c.BodyParser(&createTaskDto)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, "Information is malformed"))
	}

	userId := c.Params("id")
	savedTask := taskController.service.CreateTask(createTaskDto, userId)

	return common.SuccessResponse(c, savedTask, fiber.StatusCreated)
}

func (taskController *Controller) DeleteTask(c *fiber.Ctx) error {
	userId := c.Params("userId")
	taskId := c.Params("taskId")

	taskController.service.DeleteTaskById(taskId, userId)
	return common.SuccessResponse(c, nil, fiber.StatusNoContent)
}
