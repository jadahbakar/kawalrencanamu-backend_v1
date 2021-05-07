package main

import (
	"log"

	"github.com/gofiber/fiber/v2" // new
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/middleware"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/utils"
)

func main() {
	// Load env.
	confApp, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Define Fiber config.
	confFiber := config.FiberConfig(confApp)
	app := fiber.New(confFiber)
	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app, confApp)
}
