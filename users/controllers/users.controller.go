package UsersControllers

import (
	"github.com/HenCor2019/task-go/common/responses"
	"github.com/HenCor2019/task-go/users/dtos"
	"github.com/HenCor2019/task-go/users/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
  createUserDto := dtos.CreateUserDto{}
  c.BodyParser(&createUserDto)
  savedUser := UsersServices.CreateUser(createUserDto)
  return common.SuccessResponse(c, savedUser,fiber.StatusCreated)
}

func Find(c *fiber.Ctx) error {
  users := UsersServices.Find()
  return common.SuccessResponse(c, users,fiber.StatusOK)
}

func FindById(c *fiber.Ctx) error {
  userId := c.Params("id")
  user := UsersServices.FindById(userId)
  return common.SuccessResponse(c, user,fiber.StatusOK)
}

func DeleteById(c *fiber.Ctx) error {
  userId := c.Params("id")
  UsersServices.DeleteById(userId)
  return common.SuccessResponse(c, nil,fiber.StatusNoContent)
}
