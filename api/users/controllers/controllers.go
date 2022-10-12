package UsersControllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HenCor2019/task-go/api/users/services"
)

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
}

type Controller struct {
	service UsersServices.UserService
}

func New(service UsersServices.UserService) UserController {
	return &Controller{
		service: service,
	}
}
