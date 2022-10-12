package UsersControllers

import (
	common "github.com/HenCor2019/task-go/api/common/responses"
	"github.com/HenCor2019/task-go/api/users/dtos"
	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) CreateUser(c *fiber.Ctx) error {
	createUserDto := dtos.CreateUserDto{}
	c.BodyParser(&createUserDto)
	savedUser := controller.service.CreateUser(createUserDto)
	return common.SuccessResponse(c, savedUser, fiber.StatusCreated)
}

func (controller *Controller) Find(c *fiber.Ctx) error {
	users := controller.service.Find()
	return common.SuccessResponse(c, users, fiber.StatusOK)
}

func (controller *Controller) FindById(c *fiber.Ctx) error {
	userId := c.Params("id")
	user := controller.service.FindById(userId)
	return common.SuccessResponse(c, user, fiber.StatusOK)
}

func (controller *Controller) DeleteById(c *fiber.Ctx) error {
	userId := c.Params("id")
	controller.service.DeleteById(userId)
	return common.SuccessResponse(c, nil, fiber.StatusNoContent)
}
