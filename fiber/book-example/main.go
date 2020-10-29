package main

import (
	"github.com/avlo/fiber/book-example/book"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/books", book.GetBooks)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
