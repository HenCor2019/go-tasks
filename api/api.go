package api

import (
	"fmt"
	"os"
	"github.com/spf13/cast"

	"github.com/HenCor2019/task-go/api/common/responses"
	"github.com/HenCor2019/task-go/api/tasks/controllers"
	"github.com/HenCor2019/task-go/api/users/controllers"

	"github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
 UserControllers UsersControllers.UserController
 TaskControllers TasksControllers.TaskController
}

func New(
  userControllers UsersControllers.UserController,
  taskControllers TasksControllers.TaskController,
) *API {
	return &API{
		UserControllers: userControllers,
		TaskControllers: taskControllers,
	}
}

func (api *API) Start(app *fiber.App) error {
  app.Use(recover.New())
  v1 := app.Group("api/v1")

  v1.Get("healthcheck", func (c *fiber.Ctx) error {
    return c.SendString("ok")
  })

  v1.Get("users", api.UserControllers.Find)
  v1.Get("users/:id<int,min(1)>", api.UserControllers.FindById)

  v1.Delete("users/:id<int,min(1)>/tasks", api.TaskControllers.DeleteTask)
  v1.Delete("users/:id<int,min(1)>", api.UserControllers.DeleteById)

  v1.Post("users/:id<int,min(1)>/tasks", api.TaskControllers.CreateTask)
  v1.Post("users", api.UserControllers.CreateUser)


  v1.Use(common.NotFoundHandler)
  PORT := os.Getenv("PORT")
  return app.Listen(fmt.Sprintf(":%s", cast.ToString(PORT)))
}
