package handlers

import "github.com/gofiber/fiber/v2"

func HandlerError(c *fiber.Ctx, s int, m string) error {
	return c.Status(s).JSON(fiber.Map{
		"success": false,
		"error":   m,
	})
}
