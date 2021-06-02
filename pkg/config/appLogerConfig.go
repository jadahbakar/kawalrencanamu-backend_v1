package config

import (
	"fmt"
	"log"
	"os"
)

func AppLoggerConfig(config Config) *os.File {
	file, err := os.OpenFile(fmt.Sprintf("%s%s", config.LogFolder, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}
