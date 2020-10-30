package book

import (
	"github.com/avlo/fiber/book-example/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
	return c.SendString("one")
}

// func GetBook(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	db := database.DBConn
// 	var book Book
// 	db.Find(&book, id)
// 	c.JSON(book)
// }

// func NewBook(c *fiber.Ctx) {
// 	db := database.DBConn
// 	var book Book
// 	book.Title = "1984"
// 	book.Author = "George Orwell"
// 	book.Rating = 5
// 	db.Create(&book)
// 	c.JSON(book)
// }

// func DeleteBook(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	db := database.DBConn

// 	var book Book
// 	db.First(&book, id)
// 	if book.Title == "" {
// 		// c.Status(500).Send("No Book Found with ID")
// 		return
// 	}
// 	db.Delete(&book)
// 	// c.Send("Book Successfully deleted")
// }
