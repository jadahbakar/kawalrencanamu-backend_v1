package me

import "github.com/gofiber/fiber/v2"

// type service struct {
// 	gamesRepository ports.GamesRepository
// 	uidGen          uidgen.UIDGen
// }

func GetMe(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Kawal Rencanamu",
	})
	// return c.JSON(c.IP())
}
