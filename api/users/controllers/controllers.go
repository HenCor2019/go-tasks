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
	service *UsersServices.Service
}

func New(service *UsersServices.Service) *Controller {
	return &Controller{
		service: service,
	}
}
