package PokemonsControllers

import (
	common "github.com/HenCor2019/task-go/common/responses"
	"github.com/HenCor2019/task-go/pokemons/dtos"
	"github.com/HenCor2019/task-go/pokemons/services"
	"github.com/gofiber/fiber/v2"
)

func FindManyById(c *fiber.Ctx) error {
  findManyPokemons := new(pokemonsDtos.FindManyPokemonsDto)
  c.BodyParser(&findManyPokemons)

  pokemons := PokemonsServices.FindManyById(*findManyPokemons)
  return common.SuccessResponse(c,pokemons,fiber.StatusOK)
}
