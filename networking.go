package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func handleHTTP() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	err := app.Listen(":8080")
	if err != nil {
		logrus.Panicln(err)
	}
}
