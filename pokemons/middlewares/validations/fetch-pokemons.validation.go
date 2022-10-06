package PokemonsValidations

import (
	"github.com/HenCor2019/task-go/pokemons/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()
type ErrorResponse struct {
    FailedField string
    Tag         string
    Value       string
}

func ValidateStruct(user pokemonsDtos.FindManyPokemonsDto) []*ErrorResponse {
    var errors []*ErrorResponse
    err := validate.Struct(user)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}

func FetchPokemonsByIds(c *fiber.Ctx) error {
  fetchPokemonsByIdsDto := new(pokemonsDtos.FindManyPokemonsDto)
  err := c.BodyParser(fetchPokemonsByIdsDto)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
  }

  errors := ValidateStruct(*fetchPokemonsByIdsDto)
  if errors != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
      "success": false,
      "content": nil,
      "message": errors,
    })
  }

  return c.Next()
}
