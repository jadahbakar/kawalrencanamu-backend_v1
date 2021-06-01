package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2" // new
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/middleware"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/routes"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/utils"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/version"
)

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
	// pgConfig, err := pgx.ParseConfig(confApp.DBURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// pgConfig.Logger = zap.NewLogger(log.New("module", "pgx"))
	// conn, err := pgx.Connect(context.Background(), confApp.DBURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

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
}
