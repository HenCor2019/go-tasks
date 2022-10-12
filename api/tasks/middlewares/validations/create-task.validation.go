package TasksValidations

import (
	"github.com/HenCor2019/task-go/api/tasks/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(task tasksDtos.CreateTaskDto) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(task)
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

func CreateTask(c *fiber.Ctx) error {
	createTaskDto := new(tasksDtos.CreateTaskDto)
	err := c.BodyParser(createTaskDto)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	errors := ValidateStruct(*createTaskDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": errors,
		})
	}

	return c.Next()
}
