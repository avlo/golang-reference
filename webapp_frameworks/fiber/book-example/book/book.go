package book

import (
	"github.com/avlo/fiber/book-example/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	dba := database.DBConn
	var books []Book
	dba.Find(&books)
	return c.JSON(books)
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
