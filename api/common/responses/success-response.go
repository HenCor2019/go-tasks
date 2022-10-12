package common

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, response any, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"content": response,
		"message": nil,
	})
}
