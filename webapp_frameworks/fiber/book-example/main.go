package main

import (
	"fmt"

	"github.com/avlo/fiber/book-example/book"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/book", book.GetBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/book/", book.NewBook)
	app.Delete("/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&book.Book{})
	fmt.Println("Table Created")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
}
