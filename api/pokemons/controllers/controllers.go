package PokemonsControllers

import (
	"github.com/HenCor2019/task-go/api/pokemons/services"
	"github.com/gofiber/fiber/v2"
)

type PokemonController interface {
	FindManyById(c *fiber.Ctx) error
}

type Controller struct {
	service PokemonsServices.PokemonService
}

func New(service PokemonsServices.PokemonService) PokemonController {
	return &Controller{service}
}
