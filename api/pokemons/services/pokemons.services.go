package PokemonsServices

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/pokemons/dtos"
	"github.com/gofiber/fiber/v2"
)

func (pokemonService *Service) FindManyById(fetchPokemonsByIds pokemonsDtos.FetchPokemonsByIdsDto) []models.Pokemon {
	pokemons, err := pokemonService.gateway.FetchPokemonsByIds(fetchPokemonsByIds.Ids)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	return pokemons
}
