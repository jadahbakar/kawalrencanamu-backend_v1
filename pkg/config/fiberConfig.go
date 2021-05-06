package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig(config Config) fiber.Config {
	readTimeoutSecondsCount := config.ServerReadTimeOut

	return fiber.Config{
		Prefork:       config.Prefork,
		CaseSensitive: config.CaseSensitive,
		ReadTimeout:   time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
