package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App, logFile *os.File) {

	ConfigLogger := logger.Config{
		Next:       nil,
		Format:     "[${time}] ${pid} | ${locals:requestid} |${ip} | ${status} - ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006/Jan/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     logFile,
	}

	ConfigCORS := cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}

	ConfigCompress := compress.Config{
		Next:  nil,
		Level: compress.LevelBestSpeed,
	}

	ConfigCSRF := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}

	a.Use(
		requestid.New(),
		// Add CORS to each route.
		cors.New(ConfigCORS),
		// Add logger.
		logger.New(ConfigLogger),
		// Add CSRF
		csrf.New(ConfigCSRF),
		// Add Compress
		compress.New(ConfigCompress),
	)
}
