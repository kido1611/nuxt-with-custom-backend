package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Use(logger.New())

	return app
}
