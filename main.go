package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"` // "available" or "checked out"
}

var books []Book
var currentID int
var availableIDs []int // List of deleted IDs available for reuse

// This handler creates the book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book) // Parse the request body JSON into a Book struct

	// Reuse an ID from availableIDs if possible, otherwise increment currentID
	if len(availableIDs) > 0 {
		book.ID = availableIDs[0]
		availableIDs = availableIDs[1:] // Remove the used ID from availableIDs
	} else {
		currentID++
		book.ID = currentID
	}

	book.Status = "available"       // New books are available by default
	books = append(books, book)     // Add the new book to the books slice
	json.NewEncoder(w).Encode(book) // Return the newly created book as JSON
}

// This handler fetches all books in the collection.
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books) // Return the books slice as JSON
}

// This handler fetches a book by its id
func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// This handler updates an existing book's details.
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&book) // Parse the request body JSON
			book.ID = id                              // Make sure the ID stays the same
			books[i] = book                           // Update the book in the books slice
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// This handler deletes a book by its id
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...) // Remove the book from the slice
			availableIDs = append(availableIDs, id)   // Add the ID to the list of available IDs
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Router for the API and Server
func main() {
	router := mux.NewRouter()

	// Define the routes and map them to handler functions
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// It will Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
