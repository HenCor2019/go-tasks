package PokemonsServices

import (
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/pokemons/dtos"
	"github.com/HenCor2019/task-go/pokemons/gateways"
	"github.com/gofiber/fiber/v2"
)

func FindManyById(fetchPokemonsByIds pokemonsDtos.FetchPokemonsByIdsDto) ([]models.Pokemon) {
  pokemons,err := PokemonsGateways.FetchPokemonsByIds(fetchPokemonsByIds.Ids)

  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
  }

  return pokemons
}
