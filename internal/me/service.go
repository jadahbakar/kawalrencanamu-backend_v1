package me

import "github.com/gofiber/fiber/v2"

func GetMe(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Kawal Rencanamu Backend System",
	})
}
