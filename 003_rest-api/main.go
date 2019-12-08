package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	id     string  `json:"id"`
	isbn   string  `json:"isbn"`
	title  string  `json:"title"`
	author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	firstName string `json:"fistName"`
	lastName  string `json:"fistName"`
	age       string `json:"age"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get Book by ID
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
