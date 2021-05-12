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
	// Define Log file
	fileLogger := config.LoggerConfig(confApp)
	// Define Fiber's config.
	confFiber := config.FiberConfig(confApp)
	app := fiber.New(confFiber)
	// Register Fiber's middleware for app.
	middleware.FiberMiddleware(app, fileLogger)

	// Routes.
	// routes.SwaggerRoute(app)          // Register a route for API Docs (Swagger).
	// routes.PublicRoutes(app, confApp) // Register a public routes for app.
	// routes.PrivateRoutes(app)         // Register a private routes for app.
	// routes.NotFoundRoute(app)         // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app, confApp, fileLogger)
}
