package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:id`
	ISBN   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:id`
}

// author struct
type Author struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
}

// books
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get req params
	params := mux.Vars(r)

	//loop over books
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

		json.NewEncoder(w).Encode(&Book{})

	}
}

// create books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	// generate id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// generate id

	for i, ele := range books {
		// delete previous book
		if ele.ID == params["id"] {
			book.ID = ele.ID
			books = append(books[:i], books[i+1:]...)
			books = append(books, book)
			break
		}
	}
	json.NewEncoder(w).Encode(book)

}

// delete books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, ele := range books {
		if ele.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	// init router
	r := mux.NewRouter()

	// db - here
	books = append(books, Book{ID: "1", ISBN: "44312", Title: "First Book", Author: &Author{FirstName: "John", LastName: "Doe"}})

	books = append(books, Book{ID: "2", ISBN: "22312", Title: "Second Book", Author: &Author{FirstName: "John", LastName: "Doe"}})

	books = append(books, Book{ID: "3", ISBN: "11312", Title: "Third Book", Author: &Author{FirstName: "Ahmad", LastName: "Ali"}})

	// end points
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
