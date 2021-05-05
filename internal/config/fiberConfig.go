package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig(config Config) fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount := config.ServerReadTimeOut

	// Return Fiber configuration.
	return fiber.Config{
		Prefork:       config.Prefork,
		CaseSensitive: config.CaseSensitive,
		ReadTimeout:   time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
