package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"Go-bookstore/pkg/models"
	"Go-bookstore/pkg/utils"
)

var NewBook models.Book

// GET /books - all books with cache
func GetBooks(w http.ResponseWriter, r *http.Request) {
	cacheKey := "books_all"

	// Check Redis cache
	cachedData, err := utils.RedisClient.Get(utils.Ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedData))
		return
	}

	// Cache miss → fetch from DB
	books, err := NewBook.GetAllBooks()
	if err != nil {
		http.Error(w, "failed to fetch books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(books)

	// Store in Redis for 10 minutes
	utils.RedisClient.Set(utils.Ctx, cacheKey, string(jsonData), 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// GET /books/{bookId} - single book with cache
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	cacheKey := "book_" + bookId

	// Redis cache
	cached, err := utils.RedisClient.Get(utils.Ctx, cacheKey).Result()
	if err == nil && cached != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cached))
		return
	}

	// Cache miss → DB
	book, db := NewBook.GetBooksById(int64(ID))
	if db != nil {
		if db.RecordNotFound() || db.Error != nil {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
	} else {
		if book == nil || book.ID == 0 {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
	}

	jsonData, _ := json.Marshal(book)

	// Save to Redis
	utils.RedisClient.Set(utils.Ctx, cacheKey, string(jsonData), 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// POST /books - create book & invalidate cache
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	if book.Name == "" || book.Author == "" || book.Publication == "" {
		http.Error(w, "name, author, and publication are required", http.StatusBadRequest)
		return
	}

	b, err := book.CreateBookHandler()
	if err != nil {
		http.Error(w, "failed to create book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Invalidate cache
	utils.RedisClient.Del(utils.Ctx, "books_all")

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// PUT /books/{bookId} - update book & invalidate cache
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	bookDetails, db := NewBook.GetBooksById(int64(ID))
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

	if db != nil {
		db.Save(bookDetails)
	}

	// Invalidate cache
	utils.RedisClient.Del(utils.Ctx, "books_all")
	utils.RedisClient.Del(utils.Ctx, "book_"+bookId)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DELETE /books/{bookId} - delete book & invalidate cache
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book := NewBook.DeleteBook(int64(ID))
	if book.ID == 0 {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	// Invalidate cache
	utils.RedisClient.Del(utils.Ctx, "books_all")
	utils.RedisClient.Del(utils.Ctx, "book_"+bookId)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
