package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2" // new
	"github.com/jackc/pgx/v4"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/middleware"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/routes"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/utils"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/version"
)

func main() {
	ver := flag.Bool("version", false, "print version information")
	v := flag.Bool("v", false, "print version information")

	flag.Parse()
	if *v || *ver {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// Load env.
	confApp, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Database Connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Define Log file
	fileLogger := config.LoggerConfig(confApp)
	// Define Fiber's config.
	confFiber := config.FiberConfig(confApp)
	app := fiber.New(confFiber)
	// Register Fiber's middleware for app.
	middleware.FiberMiddleware(app, fileLogger)

	// Routes
	// routes.SwaggerRoute(app)          // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app, confApp) // Register a public routes for app.
	// routes.PrivateRoutes(app)         // Register a private routes for app.
	// routes.NotFoundRoute(app)         // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.AppName()
	utils.StartServerWithGracefulShutdown(app, confApp, fileLogger)
	// utils.StartServer(app, confApp)
}

func chooseRepo() {
	redisURL := os.Getenv("REDIS_URL")
	repo, err := rr.NewRedisRepository(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	return repo

	// return nil
}
