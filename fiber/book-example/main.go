package main

import (
	"fmt"

	"github.com/avlo/fiber/book-example/book"
	"github.com/avlo/fiber/book-example/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/book", book.GetBooks)
	// app.Get("/book/:id", book.GetBook)
	// app.Post("/book", book.NewBook)
	// app.Delete("/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		// panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(":3000")

	defer database.DBConn.Close()
}
