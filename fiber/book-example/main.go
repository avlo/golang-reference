package main

import (
	"github.com/avlo/fiber/book-example/book"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/book", book.GetBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/book/", book.NewBook)
	app.Delete("/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
