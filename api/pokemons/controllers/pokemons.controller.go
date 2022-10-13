package PokemonsControllers

import (
	"github.com/HenCor2019/task-go/api/common/responses"
	"github.com/HenCor2019/task-go/api/pokemons/dtos"
	"github.com/gofiber/fiber/v2"
)

func (pokemonController *Controller) FindManyById(c *fiber.Ctx) error {
	fetchPokemonsByIds := new(pokemonsDtos.FetchPokemonsByIdsDto)
	c.BodyParser(&fetchPokemonsByIds)

	pokemons := pokemonController.service.FindManyById(*fetchPokemonsByIds)
	return common.SuccessResponse(c, pokemons, fiber.StatusOK)
}
