package me

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(f fiber.Router) {
	f.Get("/me", GetMe)
}

func GetMe(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kawal Rencanamu Backend by @dedy"})
}
