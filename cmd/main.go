package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2" // new
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/middleware"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/routes"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/utils"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/version"
)

// https://gist.github.com/rnyrnyrny/282fe705d6e8dc012e482582d7c8ec0b

func init() {
	ver := flag.Bool("version", false, "print version information")
	v := flag.Bool("v", false, "print version information")

	flag.Parse()
	if *v || *ver {
		fmt.Println(version.String())
		os.Exit(0)
	}
}

func main() {
	// Load env.
	confApp, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Database Connection
	pgxConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to Parse DATABASE_URL: %v\n", err)
		os.Exit(1)
	}
	pgxPool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer pgxPool.Close()

	// Define Log file
	fileLogger := config.AppLoggerConfig(confApp)
	// Define Fiber's config.
	confFiber := config.FiberConfig(confApp)
	app := fiber.New(confFiber)
	// Register Fiber's middleware for app.
	middleware.FiberMiddleware(app, fileLogger)

	// Routes
	// routes.SwaggerRoute(app)          // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app, confApp, pgxPool) // Register a public routes for app.
	// routes.PrivateRoutes(app)         // Register a private routes for app.
	// routes.NotFoundRoute(app)         // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.AppName()
	utils.StartServerWithGracefulShutdown(app, confApp, fileLogger)
}
