package PokemonsServices

import (
	"os"
	"testing"

	PokemonsGateways "github.com/HenCor2019/task-go/api/pokemons/gateways"
)

var (
	gateway             *PokemonsGateways.MockPokemonGateway
	pokemonServicesMock PokemonService
)

func TestMain(m *testing.M) {
	gateway = &PokemonsGateways.MockPokemonGateway{}
	pokemonServicesMock = New(gateway)

	code := m.Run()
	os.Exit(code)
}
