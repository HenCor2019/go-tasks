package PokemonsServices

import (
	"github.com/HenCor2019/task-go/models"
	"github.com/HenCor2019/task-go/pokemons/dtos"
	"github.com/HenCor2019/task-go/pokemons/repositories"
	"github.com/gofiber/fiber/v2"
)

func FindManyById(findManyPokemonsDto pokemonsDtos.FindManyPokemonsDto) ([]models.Pokemon) {
  pokemons,err := PokemonsRepositories.FindManyById(findManyPokemonsDto.Ids)

  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, "Cannot get pokemons"))
  }

  return pokemons
}
