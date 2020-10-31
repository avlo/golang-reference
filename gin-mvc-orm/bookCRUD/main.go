package main

import (
	"bookCRUD/models"
	"bookCRUD/routes"
)

// curl -H "Content-Type: application/json" -X POST -d '{"assignedTo":"nick","task":"123"}' http://localhost:8080/tasks
func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
