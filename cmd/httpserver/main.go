package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	"github.com/jadahbakar/kawalrencanamu-backend/internal/config"
)

func main() {
	confApp, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Define Fiber config.
	confFiber := config.FiberConfig(confApp)

	log.Println("Starting Server...")
	app := fiber.New(confFiber)
	app.Use(logger.New()) // new

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// log.Println(config.HttpPort)
	if err := app.Listen(fmt.Sprintf(":%d", confApp.HttpPort)); err != nil {
		log.Panic(err)
	}
	fmt.Println("Running cleanup tasks...")
}
