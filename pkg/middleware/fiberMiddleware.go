package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App, logFile *os.File) {
	// file, err := os.OpenFile("./log/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }

	var ConfigLogger = logger.Config{
		Next:       nil,
		Format:     "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006/Jan/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     logFile,
	}

	a.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			Next:             nil,
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,PUT,DELETE",
			AllowHeaders:     "Origin, Content-Type, Accept",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		}),
		// Add logger.
		logger.New(ConfigLogger),
		// logger.New(logger.Config{
		// 	Format:     "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		// 	TimeFormat: "2006/Jan/02 15:04:05",
		// 	TimeZone:   "Asia/Jakarta",
		// 	Output:     file,
		// }),
		// Add Compress

		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
	)
}
