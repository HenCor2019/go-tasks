// Code generated by mockery v2.14.0. DO NOT EDIT.

package PokemonsGateways

import (
	models "github.com/HenCor2019/task-go/api/models"
	mock "github.com/stretchr/testify/mock"
)

// MockPokemonGateway is an autogenerated mock type for the PokemonGateway type
type MockPokemonGateway struct {
	mock.Mock
}

// FetchPokemonsByIds provides a mock function with given fields: pokemonsIds
func (_m *MockPokemonGateway) FetchPokemonsByIds(pokemonsIds []int) ([]models.Pokemon, error) {
	ret := _m.Called(pokemonsIds)

	var r0 []models.Pokemon
	if rf, ok := ret.Get(0).(func([]int) []models.Pokemon); ok {
		r0 = rf(pokemonsIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(pokemonsIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockPokemonGateway interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPokemonGateway creates a new instance of MockPokemonGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPokemonGateway(t mockConstructorTestingTNewMockPokemonGateway) *MockPokemonGateway {
	mock := &MockPokemonGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}