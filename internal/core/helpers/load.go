package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const tagName = "env"

// Load loads the environment variables into the struct
func Load() error {
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}
	return nil
}
