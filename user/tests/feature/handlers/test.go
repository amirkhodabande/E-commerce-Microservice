package handlers

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/ecommerce/user/presentation"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setup(_ *testing.T) (*fiber.App, *persistence.Container) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	err = godotenv.Load(filepath.Join(dir, "/../../../.env.testing"))
	if err != nil {
		log.Fatalf("Error loading .env.testing file")
	}

	container := persistence.NewContainer()

	app := fiber.New(fiber.Config{})

	presentation.RegisterRoutes(app, container)

	return app, container
}
