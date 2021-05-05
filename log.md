
func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Println("Starting Server...")
	app := fiber.New()

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
	if err := app.Listen(fmt.Sprintf(":%d", config.HttpPort)); err != nil {
		log.Panic(err)
	}
	fmt.Println("Running cleanup tasks...")
}
