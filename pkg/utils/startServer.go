package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App, config config.Config, logFile *os.File) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
		log.Println("Shutting down server...")
		// Clossing the Fiber Log File
		defer logFile.Close()
		log.Println("Close log file...")
	}()

	// Run server.
	if err := a.Listen(fmt.Sprintf(":%d", config.HttpPort)); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App, config config.Config) {
	// Run server.
	if err := a.Listen(fmt.Sprintf(":%d", config.HttpPort)); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
