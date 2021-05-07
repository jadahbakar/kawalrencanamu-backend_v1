package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			Next:             nil,
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
			AllowHeaders:     "Origin, Content-Type, Accept",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		}),
		// Add simple logger.
		logger.New(logger.Config{
			Format:     "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
			TimeFormat: "02/Jan/2006-15:04:05",
			TimeZone:   "Asia/Jakarta",
		}),
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
	)
}
