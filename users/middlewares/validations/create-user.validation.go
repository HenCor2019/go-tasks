package UsersValidations

import (
	"github.com/HenCor2019/task-go/users/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()
type ErrorResponse struct {
    FailedField string
    Tag         string
    Value       string
}

func ValidateStruct(user dtos.CreateUserDto) []*ErrorResponse {
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

func CreateUser(c *fiber.Ctx) error {
  createUserDto := new(dtos.CreateUserDto)
  err := c.BodyParser(createUserDto)
  if err != nil {
    panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
  }

  errors := ValidateStruct(*createUserDto)
  if errors != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "errors": errors,
    })
  }

  return c.Next()
}
