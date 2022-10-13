package PokemonsGateways

import (
	"github.com/HenCor2019/task-go/api/models"
)

type PokemonGateway interface {
	FetchPokemonsByIds(pokemonsIds []int) ([]models.Pokemon, error)
}

type Gateway struct{}

func New() PokemonGateway {
	return &Gateway{}
}
