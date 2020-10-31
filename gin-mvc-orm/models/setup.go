package models

import (

	// go-sql-driver/mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	// USER := "root"
	// PASS := "password"
	// HOST := "localhost"
	// PORT := "3306"
	// DBNAME := "bookCRUD"
	// URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("sqlite3", "task.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
