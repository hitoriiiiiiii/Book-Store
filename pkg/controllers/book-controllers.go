package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"Go-bookstore/pkg/models"
	"Go-bookstore/pkg/utils"
)

var NewBook models.Book

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := NewBook.GetAllBooks()
	if err != nil {
		http.Error(w, "failed to fetch books: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById returns a single book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}
	book, _ := NewBook.GetBooksById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book) // safely decode JSON

	if book.Name == "" || book.Author == "" || book.Publication == "" {
		http.Error(w, "name, author, and publication are required", http.StatusBadRequest)
		return
	}

	b, err := book.CreateBookHandler()
	if err != nil {
		http.Error(w, "failed to create book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book := NewBook.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook updates an existing book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	bookDetails, db := NewBook.GetBooksById(ID)
	if bookDetails == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save changes (fallback to slice or MySQL DB)
	if db != nil {
		db.Save(bookDetails)
	}

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
