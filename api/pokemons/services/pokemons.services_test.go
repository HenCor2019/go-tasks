package PokemonsServices

import (
	"errors"
	"testing"

	"github.com/HenCor2019/task-go/api/models"
	pokemonsDtos "github.com/HenCor2019/task-go/api/pokemons/dtos"
)

func TestFindManyById(t *testing.T) {
	testCases := []struct {
		Name          string
		Result        []models.Pokemon
		Dto           pokemonsDtos.FetchPokemonsByIdsDto
		ExpectedError error
	}{
		{
			Name:          "Should return all pokemons",
			Result:        []models.Pokemon{},
			Dto:           pokemonsDtos.FetchPokemonsByIdsDto{Ids: []int{1, 2}},
			ExpectedError: nil,
		},

		{
			Name:          "Should throw an error if cannot fetch",
			Result:        []models.Pokemon{},
			Dto:           pokemonsDtos.FetchPokemonsByIdsDto{Ids: []int{1, 2}},
			ExpectedError: nil,
		},
	}

	gateway.On("FetchPokemonsByIds", []int{1, 2}).Return([]models.Pokemon{}, nil).Once()
	gateway.On("FetchPokemonsByIds", []int{1, 2}).Return([]models.Pokemon{}, errors.New("Cannot fetch pokemons")).Once()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			gateway.Mock.Test(t)
			t.Parallel()

			defer func() { recover() }()
			pokemonServicesMock.FindManyById(tc.Dto)
		})
	}
}
