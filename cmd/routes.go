package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lechenco/go-rest/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/fact", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
