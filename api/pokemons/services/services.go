package PokemonsServices

import (
	"github.com/HenCor2019/task-go/api/models"
	pokemonsDtos "github.com/HenCor2019/task-go/api/pokemons/dtos"
	PokemonsGateways "github.com/HenCor2019/task-go/api/pokemons/gateways"
)

type PokemonService interface {
	FindManyById(fetchPokemonsByIds pokemonsDtos.FetchPokemonsByIdsDto) []models.Pokemon
}

type Service struct {
	gateway PokemonsGateways.PokemonGateway
}

func New(gateway PokemonsGateways.PokemonGateway) PokemonService {
	return &Service{gateway}
}
