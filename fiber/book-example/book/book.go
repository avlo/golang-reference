package book

import "github.com/gofiber/fiber"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single books")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New book")
}

func DeleteBooks(c *fiber.Ctx) error {
	return c.SendString("Delete books")
}
