package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikeyg42/Hex_demo/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListMoves)

	app.Post("/newMove", handlers.MakeMove)
}
