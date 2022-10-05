package main

import (
	"log"
	"github.com/HenCor2019/task-go/common/responses"
	"github.com/HenCor2019/task-go/config/db"
	"github.com/HenCor2019/task-go/users/controllers"
	"github.com/HenCor2019/task-go/tasks/controllers"
	"github.com/HenCor2019/task-go/pokemons/controllers"
	"github.com/HenCor2019/task-go/users/middlewares/validations"
	"github.com/HenCor2019/task-go/tasks/middlewares/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
  app := fiber.New(fiber.Config{ ErrorHandler: common.ErrorHandling, })
  v1 := app.Group("api/v1")
  app.Use(recover.New())

  v1.Get("healthcheck", func (c *fiber.Ctx) error {
    return c.SendString("ok")
  })

  v1.Get("users/", UsersControllers.Find)
  v1.Get("users/:id<int;min(1)>", UsersControllers.FindById)
  v1.Post("users/", UsersValidations.CreateUser ,UsersControllers.CreateUser)
  v1.Delete("users/:id<int,min(1)>", UsersControllers.DeleteById)

  v1.Post("users/:id<int,min(1)>/tasks", TasksValidations.CreateTask ,TasksController.CreateTask)
  v1.Delete("users/:userId<int,min(1)>/tasks/:taskId<int,min(1)>", TasksController.DeleteTask)

  v1.Get("pokemons/", PokemonsControllers.FindManyById)

  v1.Use(common.NotFoundHandler)
  db.DBConnection()
  log.Fatal(app.Listen(":3000"))
}
