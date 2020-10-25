package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article comment
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles comment
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	// i := <-ch
	// fmt.Fprintf(w, "Welcome to the HomePage [%v]", i) // writes to ResponseWriter
	fmt.Fprintf(w, "Welcome to the HomePage") // writes to ResponseWriter
	fmt.Println("Endpoint Hit: homePage")     // writes to console
	// ch <- i+1
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
