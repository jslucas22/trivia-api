package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jslucas22/trivia-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
