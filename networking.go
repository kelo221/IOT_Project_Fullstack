package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func handleHTTP() {

	app := fiber.New()
	app.Static("/", "./public")

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥈 Second handler")
		return c.Next()
	})

	// GET /api/list
	app.Get("/api/list", func(c *fiber.Ctx) error {
		fmt.Println("🥉 Last handler")
		return c.SendString("Hello, World 👋!")
	})

	err := app.Listen(":8080")
	if err != nil {
		logrus.Panicln(err)
	}
}
