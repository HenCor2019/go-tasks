package main

import (
	"context"

	"github.com/HenCor2019/task-go/api"
	"github.com/HenCor2019/task-go/api/config/db"

	"github.com/HenCor2019/task-go/api/users/controllers"
	"github.com/HenCor2019/task-go/api/users/repositories"
	"github.com/HenCor2019/task-go/api/users/services"

	"github.com/HenCor2019/task-go/api/tasks/controllers"
	"github.com/HenCor2019/task-go/api/tasks/repositories"
	"github.com/HenCor2019/task-go/api/tasks/services"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(db.New),

		fx.Provide(
			UsersRepositories.New,
			UsersServices.New,
   UsersControllers.New,
		),

		fx.Provide(
			TasksRepositories.New,
			TasksServices.New,
   TasksControllers.New,
		),

		fx.Provide(
			api.New,
			fiber.New,
		),

		fx.Invoke(setLifeCycle),
	)

	app.Run()
}

func setLifeCycle(
  lc fx.Lifecycle,
  a *api.API,
  app *fiber.App,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go a.Start(app)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
