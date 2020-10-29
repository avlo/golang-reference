package book

import "github.com/gofiber/fiber"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete book")
}
