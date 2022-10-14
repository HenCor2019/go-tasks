package common

import "github.com/gofiber/fiber/v2"

func ErrorHandling(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	err = ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"content": nil,
		"message": err.Error(),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}
