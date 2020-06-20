package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	id := strconv.Itoa(rand.Intn(100000))

	book.ID = id
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	router := mux.NewRouter()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	books = append(books, Book{ID: "1", Isbn: "3121", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "1241", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
